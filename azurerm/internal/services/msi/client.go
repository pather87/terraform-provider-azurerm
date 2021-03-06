package msi

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	UserAssignedIdentitiesClient msi.UserAssignedIdentitiesClient
}

func BuildClient(o *common.ClientOptions) *Client {
	c := Client{}

	c.UserAssignedIdentitiesClient = msi.NewUserAssignedIdentitiesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&c.UserAssignedIdentitiesClient.Client, o.ResourceManagerAuthorizer)

	return &c
}
