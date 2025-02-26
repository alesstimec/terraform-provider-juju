// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/terraform-provider-juju/internal/juju (interfaces: SharedClient,ClientAPIClient,ApplicationAPIClient,ModelConfigAPIClient,ResourceAPIClient,SecretAPIClient,JaasAPIClient,KubernetesCloudAPIClient)
//
// Generated by this command:
//
//	mockgen -package juju -destination mock_test.go github.com/juju/terraform-provider-juju/internal/juju SharedClient,ClientAPIClient,ApplicationAPIClient,ModelConfigAPIClient,ResourceAPIClient,SecretAPIClient,JaasAPIClient,KubernetesCloudAPIClient
//

// Package juju is a generated GoMock package.
package juju

import (
	io "io"
	reflect "reflect"

	params "github.com/canonical/jimm-go-sdk/v3/api/params"
	charm "github.com/juju/charm/v12"
	resource "github.com/juju/charm/v12/resource"
	api "github.com/juju/juju/api"
	application "github.com/juju/juju/api/client/application"
	client "github.com/juju/juju/api/client/client"
	resources "github.com/juju/juju/api/client/resources"
	secrets "github.com/juju/juju/api/client/secrets"
	charm0 "github.com/juju/juju/api/common/charm"
	cloud "github.com/juju/juju/cloud"
	constraints "github.com/juju/juju/core/constraints"
	model "github.com/juju/juju/core/model"
	resources0 "github.com/juju/juju/core/resources"
	secrets0 "github.com/juju/juju/core/secrets"
	params0 "github.com/juju/juju/rpc/params"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockSharedClient is a mock of SharedClient interface.
type MockSharedClient struct {
	ctrl     *gomock.Controller
	recorder *MockSharedClientMockRecorder
	isgomock struct{}
}

// MockSharedClientMockRecorder is the mock recorder for MockSharedClient.
type MockSharedClientMockRecorder struct {
	mock *MockSharedClient
}

// NewMockSharedClient creates a new mock instance.
func NewMockSharedClient(ctrl *gomock.Controller) *MockSharedClient {
	mock := &MockSharedClient{ctrl: ctrl}
	mock.recorder = &MockSharedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedClient) EXPECT() *MockSharedClientMockRecorder {
	return m.recorder
}

// AddModel mocks base method.
func (m *MockSharedClient) AddModel(modelName, modelUUID string, modelType model.ModelType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddModel", modelName, modelUUID, modelType)
}

// AddModel indicates an expected call of AddModel.
func (mr *MockSharedClientMockRecorder) AddModel(modelName, modelUUID, modelType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddModel", reflect.TypeOf((*MockSharedClient)(nil).AddModel), modelName, modelUUID, modelType)
}

// Debugf mocks base method.
func (m *MockSharedClient) Debugf(msg string, additionalFields ...map[string]any) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range additionalFields {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockSharedClientMockRecorder) Debugf(msg any, additionalFields ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, additionalFields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockSharedClient)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockSharedClient) Errorf(err error, msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Errorf", err, msg)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockSharedClientMockRecorder) Errorf(err, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockSharedClient)(nil).Errorf), err, msg)
}

// GetConnection mocks base method.
func (m *MockSharedClient) GetConnection(modelName *string) (api.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", modelName)
	ret0, _ := ret[0].(api.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockSharedClientMockRecorder) GetConnection(modelName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockSharedClient)(nil).GetConnection), modelName)
}

// JujuLogger mocks base method.
func (m *MockSharedClient) JujuLogger() *jujuLoggerShim {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JujuLogger")
	ret0, _ := ret[0].(*jujuLoggerShim)
	return ret0
}

// JujuLogger indicates an expected call of JujuLogger.
func (mr *MockSharedClientMockRecorder) JujuLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JujuLogger", reflect.TypeOf((*MockSharedClient)(nil).JujuLogger))
}

// ModelStatus mocks base method.
func (m *MockSharedClient) ModelStatus(modelIfentifier string, conn api.Connection) (*params0.FullStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelStatus", modelIfentifier, conn)
	ret0, _ := ret[0].(*params0.FullStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelStatus indicates an expected call of ModelStatus.
func (mr *MockSharedClientMockRecorder) ModelStatus(modelIfentifier, conn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelStatus", reflect.TypeOf((*MockSharedClient)(nil).ModelStatus), modelIfentifier, conn)
}

// ModelType mocks base method.
func (m *MockSharedClient) ModelType(modelName string) (model.ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelType", modelName)
	ret0, _ := ret[0].(model.ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelType indicates an expected call of ModelType.
func (mr *MockSharedClientMockRecorder) ModelType(modelName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelType", reflect.TypeOf((*MockSharedClient)(nil).ModelType), modelName)
}

// ModelUUID mocks base method.
func (m *MockSharedClient) ModelUUID(modelName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID", modelName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockSharedClientMockRecorder) ModelUUID(modelName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockSharedClient)(nil).ModelUUID), modelName)
}

// RemoveModel mocks base method.
func (m *MockSharedClient) RemoveModel(modelUUID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveModel", modelUUID)
}

// RemoveModel indicates an expected call of RemoveModel.
func (mr *MockSharedClientMockRecorder) RemoveModel(modelUUID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModel", reflect.TypeOf((*MockSharedClient)(nil).RemoveModel), modelUUID)
}

// Tracef mocks base method.
func (m *MockSharedClient) Tracef(msg string, additionalFields ...map[string]any) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range additionalFields {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockSharedClientMockRecorder) Tracef(msg any, additionalFields ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, additionalFields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockSharedClient)(nil).Tracef), varargs...)
}

// Warnf mocks base method.
func (m *MockSharedClient) Warnf(msg string, additionalFields ...map[string]any) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range additionalFields {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockSharedClientMockRecorder) Warnf(msg any, additionalFields ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, additionalFields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockSharedClient)(nil).Warnf), varargs...)
}

// MockClientAPIClient is a mock of ClientAPIClient interface.
type MockClientAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientAPIClientMockRecorder
	isgomock struct{}
}

// MockClientAPIClientMockRecorder is the mock recorder for MockClientAPIClient.
type MockClientAPIClientMockRecorder struct {
	mock *MockClientAPIClient
}

// NewMockClientAPIClient creates a new mock instance.
func NewMockClientAPIClient(ctrl *gomock.Controller) *MockClientAPIClient {
	mock := &MockClientAPIClient{ctrl: ctrl}
	mock.recorder = &MockClientAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientAPIClient) EXPECT() *MockClientAPIClientMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockClientAPIClient) Status(args *client.StatusArgs) (*params0.FullStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", args)
	ret0, _ := ret[0].(*params0.FullStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockClientAPIClientMockRecorder) Status(args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockClientAPIClient)(nil).Status), args)
}

// MockApplicationAPIClient is a mock of ApplicationAPIClient interface.
type MockApplicationAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationAPIClientMockRecorder
	isgomock struct{}
}

// MockApplicationAPIClientMockRecorder is the mock recorder for MockApplicationAPIClient.
type MockApplicationAPIClientMockRecorder struct {
	mock *MockApplicationAPIClient
}

// NewMockApplicationAPIClient creates a new mock instance.
func NewMockApplicationAPIClient(ctrl *gomock.Controller) *MockApplicationAPIClient {
	mock := &MockApplicationAPIClient{ctrl: ctrl}
	mock.recorder = &MockApplicationAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationAPIClient) EXPECT() *MockApplicationAPIClientMockRecorder {
	return m.recorder
}

// AddUnits mocks base method.
func (m *MockApplicationAPIClient) AddUnits(args application.AddUnitsParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUnits", args)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUnits indicates an expected call of AddUnits.
func (mr *MockApplicationAPIClientMockRecorder) AddUnits(args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnits", reflect.TypeOf((*MockApplicationAPIClient)(nil).AddUnits), args)
}

// ApplicationsInfo mocks base method.
func (m *MockApplicationAPIClient) ApplicationsInfo(applications []names.ApplicationTag) ([]params0.ApplicationInfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationsInfo", applications)
	ret0, _ := ret[0].([]params0.ApplicationInfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationsInfo indicates an expected call of ApplicationsInfo.
func (mr *MockApplicationAPIClientMockRecorder) ApplicationsInfo(applications any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationsInfo", reflect.TypeOf((*MockApplicationAPIClient)(nil).ApplicationsInfo), applications)
}

// Deploy mocks base method.
func (m *MockApplicationAPIClient) Deploy(args application.DeployArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", args)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockApplicationAPIClientMockRecorder) Deploy(args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockApplicationAPIClient)(nil).Deploy), args)
}

// DeployFromRepository mocks base method.
func (m *MockApplicationAPIClient) DeployFromRepository(arg application.DeployFromRepositoryArg) (application.DeployInfo, []application.PendingResourceUpload, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployFromRepository", arg)
	ret0, _ := ret[0].(application.DeployInfo)
	ret1, _ := ret[1].([]application.PendingResourceUpload)
	ret2, _ := ret[2].([]error)
	return ret0, ret1, ret2
}

// DeployFromRepository indicates an expected call of DeployFromRepository.
func (mr *MockApplicationAPIClientMockRecorder) DeployFromRepository(arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployFromRepository", reflect.TypeOf((*MockApplicationAPIClient)(nil).DeployFromRepository), arg)
}

// DestroyApplications mocks base method.
func (m *MockApplicationAPIClient) DestroyApplications(in application.DestroyApplicationsParams) ([]params0.DestroyApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyApplications", in)
	ret0, _ := ret[0].([]params0.DestroyApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DestroyApplications indicates an expected call of DestroyApplications.
func (mr *MockApplicationAPIClientMockRecorder) DestroyApplications(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyApplications", reflect.TypeOf((*MockApplicationAPIClient)(nil).DestroyApplications), in)
}

// DestroyUnits mocks base method.
func (m *MockApplicationAPIClient) DestroyUnits(in application.DestroyUnitsParams) ([]params0.DestroyUnitResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyUnits", in)
	ret0, _ := ret[0].([]params0.DestroyUnitResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DestroyUnits indicates an expected call of DestroyUnits.
func (mr *MockApplicationAPIClientMockRecorder) DestroyUnits(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyUnits", reflect.TypeOf((*MockApplicationAPIClient)(nil).DestroyUnits), in)
}

// Expose mocks base method.
func (m *MockApplicationAPIClient) Expose(application string, exposedEndpoints map[string]params0.ExposedEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expose", application, exposedEndpoints)
	ret0, _ := ret[0].(error)
	return ret0
}

// Expose indicates an expected call of Expose.
func (mr *MockApplicationAPIClientMockRecorder) Expose(application, exposedEndpoints any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expose", reflect.TypeOf((*MockApplicationAPIClient)(nil).Expose), application, exposedEndpoints)
}

// Get mocks base method.
func (m *MockApplicationAPIClient) Get(branchName, application string) (*params0.ApplicationGetResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", branchName, application)
	ret0, _ := ret[0].(*params0.ApplicationGetResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApplicationAPIClientMockRecorder) Get(branchName, application any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApplicationAPIClient)(nil).Get), branchName, application)
}

// GetCharmURLOrigin mocks base method.
func (m *MockApplicationAPIClient) GetCharmURLOrigin(branchName, applicationName string) (*charm.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmURLOrigin", branchName, applicationName)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCharmURLOrigin indicates an expected call of GetCharmURLOrigin.
func (mr *MockApplicationAPIClientMockRecorder) GetCharmURLOrigin(branchName, applicationName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmURLOrigin", reflect.TypeOf((*MockApplicationAPIClient)(nil).GetCharmURLOrigin), branchName, applicationName)
}

// GetConstraints mocks base method.
func (m *MockApplicationAPIClient) GetConstraints(applications ...string) ([]constraints.Value, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range applications {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConstraints", varargs...)
	ret0, _ := ret[0].([]constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConstraints indicates an expected call of GetConstraints.
func (mr *MockApplicationAPIClientMockRecorder) GetConstraints(applications ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConstraints", reflect.TypeOf((*MockApplicationAPIClient)(nil).GetConstraints), applications...)
}

// MergeBindings mocks base method.
func (m *MockApplicationAPIClient) MergeBindings(req params0.ApplicationMergeBindingsArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeBindings", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// MergeBindings indicates an expected call of MergeBindings.
func (mr *MockApplicationAPIClientMockRecorder) MergeBindings(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeBindings", reflect.TypeOf((*MockApplicationAPIClient)(nil).MergeBindings), req)
}

// ScaleApplication mocks base method.
func (m *MockApplicationAPIClient) ScaleApplication(in application.ScaleApplicationParams) (params0.ScaleApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleApplication", in)
	ret0, _ := ret[0].(params0.ScaleApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScaleApplication indicates an expected call of ScaleApplication.
func (mr *MockApplicationAPIClientMockRecorder) ScaleApplication(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleApplication", reflect.TypeOf((*MockApplicationAPIClient)(nil).ScaleApplication), in)
}

// SetCharm mocks base method.
func (m *MockApplicationAPIClient) SetCharm(branchName string, cfg application.SetCharmConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", branchName, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharm indicates an expected call of SetCharm.
func (mr *MockApplicationAPIClientMockRecorder) SetCharm(branchName, cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockApplicationAPIClient)(nil).SetCharm), branchName, cfg)
}

// SetConfig mocks base method.
func (m *MockApplicationAPIClient) SetConfig(branchName, application, configYAML string, config map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", branchName, application, configYAML, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockApplicationAPIClientMockRecorder) SetConfig(branchName, application, configYAML, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockApplicationAPIClient)(nil).SetConfig), branchName, application, configYAML, config)
}

// SetConstraints mocks base method.
func (m *MockApplicationAPIClient) SetConstraints(application string, constraints constraints.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConstraints", application, constraints)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints.
func (mr *MockApplicationAPIClientMockRecorder) SetConstraints(application, constraints any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockApplicationAPIClient)(nil).SetConstraints), application, constraints)
}

// Unexpose mocks base method.
func (m *MockApplicationAPIClient) Unexpose(application string, endpoints []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unexpose", application, endpoints)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unexpose indicates an expected call of Unexpose.
func (mr *MockApplicationAPIClientMockRecorder) Unexpose(application, endpoints any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unexpose", reflect.TypeOf((*MockApplicationAPIClient)(nil).Unexpose), application, endpoints)
}

// MockModelConfigAPIClient is a mock of ModelConfigAPIClient interface.
type MockModelConfigAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigAPIClientMockRecorder
	isgomock struct{}
}

// MockModelConfigAPIClientMockRecorder is the mock recorder for MockModelConfigAPIClient.
type MockModelConfigAPIClientMockRecorder struct {
	mock *MockModelConfigAPIClient
}

// NewMockModelConfigAPIClient creates a new mock instance.
func NewMockModelConfigAPIClient(ctrl *gomock.Controller) *MockModelConfigAPIClient {
	mock := &MockModelConfigAPIClient{ctrl: ctrl}
	mock.recorder = &MockModelConfigAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigAPIClient) EXPECT() *MockModelConfigAPIClientMockRecorder {
	return m.recorder
}

// ModelGet mocks base method.
func (m *MockModelConfigAPIClient) ModelGet() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet.
func (mr *MockModelConfigAPIClientMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockModelConfigAPIClient)(nil).ModelGet))
}

// MockResourceAPIClient is a mock of ResourceAPIClient interface.
type MockResourceAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourceAPIClientMockRecorder
	isgomock struct{}
}

// MockResourceAPIClientMockRecorder is the mock recorder for MockResourceAPIClient.
type MockResourceAPIClientMockRecorder struct {
	mock *MockResourceAPIClient
}

// NewMockResourceAPIClient creates a new mock instance.
func NewMockResourceAPIClient(ctrl *gomock.Controller) *MockResourceAPIClient {
	mock := &MockResourceAPIClient{ctrl: ctrl}
	mock.recorder = &MockResourceAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceAPIClient) EXPECT() *MockResourceAPIClientMockRecorder {
	return m.recorder
}

// AddPendingResources mocks base method.
func (m *MockResourceAPIClient) AddPendingResources(args resources.AddPendingResourcesArgs) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPendingResources", args)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPendingResources indicates an expected call of AddPendingResources.
func (mr *MockResourceAPIClientMockRecorder) AddPendingResources(args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPendingResources", reflect.TypeOf((*MockResourceAPIClient)(nil).AddPendingResources), args)
}

// ListResources mocks base method.
func (m *MockResourceAPIClient) ListResources(applications []string) ([]resources0.ApplicationResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", applications)
	ret0, _ := ret[0].([]resources0.ApplicationResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockResourceAPIClientMockRecorder) ListResources(applications any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockResourceAPIClient)(nil).ListResources), applications)
}

// Upload mocks base method.
func (m *MockResourceAPIClient) Upload(application, name, filename, pendingID string, reader io.ReadSeeker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", application, name, filename, pendingID, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockResourceAPIClientMockRecorder) Upload(application, name, filename, pendingID, reader any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockResourceAPIClient)(nil).Upload), application, name, filename, pendingID, reader)
}

// UploadPendingResource mocks base method.
func (m *MockResourceAPIClient) UploadPendingResource(applicationID string, resource resource.Resource, filename string, r io.ReadSeeker) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPendingResource", applicationID, resource, filename, r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPendingResource indicates an expected call of UploadPendingResource.
func (mr *MockResourceAPIClientMockRecorder) UploadPendingResource(applicationID, resource, filename, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPendingResource", reflect.TypeOf((*MockResourceAPIClient)(nil).UploadPendingResource), applicationID, resource, filename, r)
}

// MockSecretAPIClient is a mock of SecretAPIClient interface.
type MockSecretAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretAPIClientMockRecorder
	isgomock struct{}
}

// MockSecretAPIClientMockRecorder is the mock recorder for MockSecretAPIClient.
type MockSecretAPIClientMockRecorder struct {
	mock *MockSecretAPIClient
}

// NewMockSecretAPIClient creates a new mock instance.
func NewMockSecretAPIClient(ctrl *gomock.Controller) *MockSecretAPIClient {
	mock := &MockSecretAPIClient{ctrl: ctrl}
	mock.recorder = &MockSecretAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretAPIClient) EXPECT() *MockSecretAPIClientMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockSecretAPIClient) CreateSecret(name, description string, data map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", name, description, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockSecretAPIClientMockRecorder) CreateSecret(name, description, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretAPIClient)(nil).CreateSecret), name, description, data)
}

// GrantSecret mocks base method.
func (m *MockSecretAPIClient) GrantSecret(uri *secrets0.URI, name string, apps []string) ([]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantSecret", uri, name, apps)
	ret0, _ := ret[0].([]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrantSecret indicates an expected call of GrantSecret.
func (mr *MockSecretAPIClientMockRecorder) GrantSecret(uri, name, apps any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantSecret", reflect.TypeOf((*MockSecretAPIClient)(nil).GrantSecret), uri, name, apps)
}

// ListSecrets mocks base method.
func (m *MockSecretAPIClient) ListSecrets(reveal bool, filter secrets0.Filter) ([]secrets.SecretDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", reveal, filter)
	ret0, _ := ret[0].([]secrets.SecretDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretAPIClientMockRecorder) ListSecrets(reveal, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretAPIClient)(nil).ListSecrets), reveal, filter)
}

// RemoveSecret mocks base method.
func (m *MockSecretAPIClient) RemoveSecret(uri *secrets0.URI, name string, revision *int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecret", uri, name, revision)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecret indicates an expected call of RemoveSecret.
func (mr *MockSecretAPIClientMockRecorder) RemoveSecret(uri, name, revision any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecret", reflect.TypeOf((*MockSecretAPIClient)(nil).RemoveSecret), uri, name, revision)
}

// RevokeSecret mocks base method.
func (m *MockSecretAPIClient) RevokeSecret(uri *secrets0.URI, name string, apps []string) ([]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecret", uri, name, apps)
	ret0, _ := ret[0].([]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecret indicates an expected call of RevokeSecret.
func (mr *MockSecretAPIClientMockRecorder) RevokeSecret(uri, name, apps any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecret", reflect.TypeOf((*MockSecretAPIClient)(nil).RevokeSecret), uri, name, apps)
}

// UpdateSecret mocks base method.
func (m *MockSecretAPIClient) UpdateSecret(uri *secrets0.URI, name string, autoPrune *bool, newName, description string, data map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", uri, name, autoPrune, newName, description, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockSecretAPIClientMockRecorder) UpdateSecret(uri, name, autoPrune, newName, description, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretAPIClient)(nil).UpdateSecret), uri, name, autoPrune, newName, description, data)
}

// MockJaasAPIClient is a mock of JaasAPIClient interface.
type MockJaasAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockJaasAPIClientMockRecorder
	isgomock struct{}
}

// MockJaasAPIClientMockRecorder is the mock recorder for MockJaasAPIClient.
type MockJaasAPIClientMockRecorder struct {
	mock *MockJaasAPIClient
}

// NewMockJaasAPIClient creates a new mock instance.
func NewMockJaasAPIClient(ctrl *gomock.Controller) *MockJaasAPIClient {
	mock := &MockJaasAPIClient{ctrl: ctrl}
	mock.recorder = &MockJaasAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJaasAPIClient) EXPECT() *MockJaasAPIClientMockRecorder {
	return m.recorder
}

// AddGroup mocks base method.
func (m *MockJaasAPIClient) AddGroup(req *params.AddGroupRequest) (params.AddGroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGroup", req)
	ret0, _ := ret[0].(params.AddGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGroup indicates an expected call of AddGroup.
func (mr *MockJaasAPIClientMockRecorder) AddGroup(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroup", reflect.TypeOf((*MockJaasAPIClient)(nil).AddGroup), req)
}

// AddRelation mocks base method.
func (m *MockJaasAPIClient) AddRelation(req *params.AddRelationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRelation", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRelation indicates an expected call of AddRelation.
func (mr *MockJaasAPIClientMockRecorder) AddRelation(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRelation", reflect.TypeOf((*MockJaasAPIClient)(nil).AddRelation), req)
}

// GetGroup mocks base method.
func (m *MockJaasAPIClient) GetGroup(req *params.GetGroupRequest) (params.GetGroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", req)
	ret0, _ := ret[0].(params.GetGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockJaasAPIClientMockRecorder) GetGroup(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockJaasAPIClient)(nil).GetGroup), req)
}

// ListRelationshipTuples mocks base method.
func (m *MockJaasAPIClient) ListRelationshipTuples(req *params.ListRelationshipTuplesRequest) (*params.ListRelationshipTuplesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRelationshipTuples", req)
	ret0, _ := ret[0].(*params.ListRelationshipTuplesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRelationshipTuples indicates an expected call of ListRelationshipTuples.
func (mr *MockJaasAPIClientMockRecorder) ListRelationshipTuples(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRelationshipTuples", reflect.TypeOf((*MockJaasAPIClient)(nil).ListRelationshipTuples), req)
}

// RemoveGroup mocks base method.
func (m *MockJaasAPIClient) RemoveGroup(req *params.RemoveGroupRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGroup", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveGroup indicates an expected call of RemoveGroup.
func (mr *MockJaasAPIClientMockRecorder) RemoveGroup(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroup", reflect.TypeOf((*MockJaasAPIClient)(nil).RemoveGroup), req)
}

// RemoveRelation mocks base method.
func (m *MockJaasAPIClient) RemoveRelation(req *params.RemoveRelationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRelation", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRelation indicates an expected call of RemoveRelation.
func (mr *MockJaasAPIClientMockRecorder) RemoveRelation(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRelation", reflect.TypeOf((*MockJaasAPIClient)(nil).RemoveRelation), req)
}

// RenameGroup mocks base method.
func (m *MockJaasAPIClient) RenameGroup(req *params.RenameGroupRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameGroup", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameGroup indicates an expected call of RenameGroup.
func (mr *MockJaasAPIClientMockRecorder) RenameGroup(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameGroup", reflect.TypeOf((*MockJaasAPIClient)(nil).RenameGroup), req)
}

// MockKubernetesCloudAPIClient is a mock of KubernetesCloudAPIClient interface.
type MockKubernetesCloudAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesCloudAPIClientMockRecorder
	isgomock struct{}
}

// MockKubernetesCloudAPIClientMockRecorder is the mock recorder for MockKubernetesCloudAPIClient.
type MockKubernetesCloudAPIClientMockRecorder struct {
	mock *MockKubernetesCloudAPIClient
}

// NewMockKubernetesCloudAPIClient creates a new mock instance.
func NewMockKubernetesCloudAPIClient(ctrl *gomock.Controller) *MockKubernetesCloudAPIClient {
	mock := &MockKubernetesCloudAPIClient{ctrl: ctrl}
	mock.recorder = &MockKubernetesCloudAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesCloudAPIClient) EXPECT() *MockKubernetesCloudAPIClientMockRecorder {
	return m.recorder
}

// AddCloud mocks base method.
func (m *MockKubernetesCloudAPIClient) AddCloud(cloud cloud.Cloud, force bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCloud", cloud, force)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCloud indicates an expected call of AddCloud.
func (mr *MockKubernetesCloudAPIClientMockRecorder) AddCloud(cloud, force any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCloud", reflect.TypeOf((*MockKubernetesCloudAPIClient)(nil).AddCloud), cloud, force)
}

// AddCredential mocks base method.
func (m *MockKubernetesCloudAPIClient) AddCredential(cloud string, credential cloud.Credential) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredential", cloud, credential)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCredential indicates an expected call of AddCredential.
func (mr *MockKubernetesCloudAPIClientMockRecorder) AddCredential(cloud, credential any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredential", reflect.TypeOf((*MockKubernetesCloudAPIClient)(nil).AddCredential), cloud, credential)
}

// Cloud mocks base method.
func (m *MockKubernetesCloudAPIClient) Cloud(tag names.CloudTag) (cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud", tag)
	ret0, _ := ret[0].(cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockKubernetesCloudAPIClientMockRecorder) Cloud(tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockKubernetesCloudAPIClient)(nil).Cloud), tag)
}

// RemoveCloud mocks base method.
func (m *MockKubernetesCloudAPIClient) RemoveCloud(cloud string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCloud", cloud)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCloud indicates an expected call of RemoveCloud.
func (mr *MockKubernetesCloudAPIClientMockRecorder) RemoveCloud(cloud any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCloud", reflect.TypeOf((*MockKubernetesCloudAPIClient)(nil).RemoveCloud), cloud)
}

// UpdateCloud mocks base method.
func (m *MockKubernetesCloudAPIClient) UpdateCloud(cloud cloud.Cloud) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCloud", cloud)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCloud indicates an expected call of UpdateCloud.
func (mr *MockKubernetesCloudAPIClientMockRecorder) UpdateCloud(cloud any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCloud", reflect.TypeOf((*MockKubernetesCloudAPIClient)(nil).UpdateCloud), cloud)
}

// UserCredentials mocks base method.
func (m *MockKubernetesCloudAPIClient) UserCredentials(user names.UserTag, cloud names.CloudTag) ([]names.CloudCredentialTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCredentials", user, cloud)
	ret0, _ := ret[0].([]names.CloudCredentialTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCredentials indicates an expected call of UserCredentials.
func (mr *MockKubernetesCloudAPIClientMockRecorder) UserCredentials(user, cloud any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCredentials", reflect.TypeOf((*MockKubernetesCloudAPIClient)(nil).UserCredentials), user, cloud)
}
