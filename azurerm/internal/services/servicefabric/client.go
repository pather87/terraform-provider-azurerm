package servicefabric

import (
	"github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2018-02-01/servicefabric"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	ClustersClient servicefabric.ClustersClient
}

func BuildClient(o *common.ClientOptions) *Client {
	c := Client{}

	c.ClustersClient = servicefabric.NewClustersClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&c.ClustersClient.Client, o.ResourceManagerAuthorizer)

	return &c
}
