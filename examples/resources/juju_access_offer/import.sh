# Access Offers can be imported by using the Offer URL as in the juju show-offers output.
# Example:
# $juju show-offer mysql
# Store            URL             Access  Description                                    Endpoint  Interface  Role
# mycontroller     admin/db.mysql  admin   MariaDB Server is one of the most ...          mysql     mysql      provider
$ terraform import juju_access_offer.db admin/db.mysql
