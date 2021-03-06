package logic

import (
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2016-06-01/logic"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	WorkflowsClient logic.WorkflowsClient
}

func BuildClient(o *common.ClientOptions) *Client {
	c := Client{}

	c.WorkflowsClient = logic.NewWorkflowsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&c.WorkflowsClient.Client, o.ResourceManagerAuthorizer)

	return &c
}
