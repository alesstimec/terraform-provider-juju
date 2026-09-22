// Copyright 2026 Canonical Ltd.
// Licensed under the Apache License, Version 2.0, see LICENCE file for details.

package juju

import (
	"context"
	"fmt"

	"github.com/juju/juju/api"
	"github.com/juju/juju/api/client/application"
	"github.com/juju/juju/api/client/client"
	"github.com/juju/juju/api/client/highavailability"
	"github.com/juju/juju/api/client/modelconfig"
	"github.com/juju/juju/api/connector"
	"github.com/juju/juju/core/constraints"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/environs/config"
)

// EnableHAInput contains the input for enabling high availability on a controller.
type EnableHAInput struct {
	// ConnInfo holds the connection details for the target controller.
	ConnInfo ControllerConnectionInformation
	// Constraints is an optional constraint string for newly provisioned
	// controller units (e.g. "mem=8G cores=4").
	Constraints string
	// Units is the desired number of controller units. On Juju < 4.0 this
	// must be odd and >= 3.
	Units int
	// To is an optional list of placement directives for the new controller
	// units (e.g. ["lxd:0", "lxd:1"]). When empty, Juju selects placement
	// automatically.
	To []string
}

// EnableHAClient handles high-availability operations against a Juju controller.
// It creates its own API connections using the connection details provided per
// invocation, keeping the client stateless and safe for concurrent use.
type EnableHAClient struct{}

// NewEnableHAClient returns a new EnableHAClient.
func NewEnableHAClient() *EnableHAClient {
	return &EnableHAClient{}
}

// EnableHA scales the controller described by input.ConnInfo up to
// input.Units units. It is idempotent; scaling down is not supported and
// must be done via the Juju CLI ("juju remove-unit").
//
// On Juju < 4.0 it uses the HighAvailability facade; on Juju >= 4.0, where
// the facade was removed, it adds units to the "controller" application
// instead, mirroring "juju add-unit -m controller controller -n N".
func (c *EnableHAClient) EnableHA(ctx context.Context, input EnableHAInput) error {
	modelUUID, err := c.controllerModelUUID(ctx, input.ConnInfo)
	if err != nil {
		return err
	}

	conn, err := c.connect(ctx, input.ConnInfo, modelUUID)
	if err != nil {
		return err
	}
	defer conn.Close()

	version, ok := conn.ServerVersion()
	if !ok {
		return fmt.Errorf("failed to get controller version")
	}
	if version.Major >= 4 {
		return c.enableHAByScaling(ctx, conn, input)
	}
	return c.enableHAViaFacade(ctx, conn, input)
}

// enableHAViaFacade is the Juju < 4.0 implementation. The facade
// reconciles the controller towards the desired unit count itself,
// promoting existing machines to voters before provisioning new ones.
func (c *EnableHAClient) enableHAViaFacade(ctx context.Context, conn api.Connection, input EnableHAInput) error {
	var cons constraints.Value
	if input.Constraints != "" {
		parsed, err := constraints.Parse(input.Constraints)
		if err != nil {
			return fmt.Errorf("failed to parse constraints %q: %w", input.Constraints, err)
		}
		cons = parsed
	}

	haClient := highavailability.NewClient(conn)
	defer haClient.Close()
	_, err := haClient.EnableHA(ctx, input.Units, cons, input.To)
	return err
}

// enableHAByScaling is the Juju >= 4.0 implementation: it adds units to
// the "controller" application in the controller model.
func (c *EnableHAClient) enableHAByScaling(ctx context.Context, conn api.Connection, input EnableHAInput) error {
	status, err := client.NewClient(conn, nil).Status(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get controller model status: %w", err)
	}
	app, ok := status.Applications["controller"]
	if !ok {
		return fmt.Errorf("controller application not found in controller model status")
	}

	diff := input.Units - len(app.Units)
	switch {
	case diff == 0:
		return nil
	case diff < 0:
		// Juju 4 could technically scale down here (the API allows
		// removing units), but we keep the Juju 3 behaviour of rejecting
		// scale-down for consistency across versions. Removing controller
		// units is an operator-guided action; use the Juju CLI
		// ("juju remove-unit") instead.
		return fmt.Errorf("removing %d controller unit(s) is not supported: use the Juju CLI (\"juju remove-unit\") instead", -diff)
	}

	placement := make([]*instance.Placement, 0, len(input.To))
	for _, directive := range input.To {
		p, err := instance.ParsePlacement(directive)
		if err != nil {
			return fmt.Errorf("failed to parse placement directive %q: %w", directive, err)
		}
		placement = append(placement, p)
	}

	if _, err := application.NewClient(conn).AddUnits(ctx, application.AddUnitsParams{
		ApplicationName: "controller",
		NumUnits:        diff,
		Placement:       placement,
	}); err != nil {
		return fmt.Errorf("failed to add %d controller units: %w", diff, err)
	}
	return nil
}

// controllerModelUUID returns the UUID of the "controller" model on the
// controller described by connInfo. A controller-only connection returns
// the controller model's own configuration, so its UUID can be read
// directly from the model config, the same way the bootstrap code does.
func (c *EnableHAClient) controllerModelUUID(ctx context.Context, connInfo ControllerConnectionInformation) (string, error) {
	conn, err := c.connect(ctx, connInfo, "")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	modelAttrs, err := modelconfig.NewClient(conn).ModelGet(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get model config: %w", err)
	}
	cfg, err := config.New(config.NoDefaults, modelAttrs)
	if err != nil {
		return "", fmt.Errorf("failed to build model config: %w", err)
	}
	if cfg.UUID() == "" {
		return "", fmt.Errorf("controller model UUID not found in model config")
	}
	return cfg.UUID(), nil
}

// connect returns a connection to the controller described by connInfo,
// scoped to the model with the given UUID (controller-only when empty).
func (c *EnableHAClient) connect(ctx context.Context, connInfo ControllerConnectionInformation, modelUUID string) (api.Connection, error) {
	connr, err := connector.NewSimple(connector.SimpleConfig{
		ControllerAddresses: connInfo.Addresses,
		CACert:              connInfo.CACert,
		Username:            connInfo.Username,
		Password:            connInfo.Password,
		ModelUUID:           modelUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create connector: %w", err)
	}
	conn, err := connr.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to controller: %w", err)
	}
	return conn, nil
}
