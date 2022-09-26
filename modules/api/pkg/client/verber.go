package client

import (
	"github.com/taidevops/dashboard/src/app/backend/api"
	restclient "k8s.io/client-go/rest"
)

type resourceVerber struct {
	client              RESTClient
	appsClient          RESTClient
	batchClient         RESTClient
	betaBatchClient     RESTClient
	autoscalingClient   RESTClient
	storageClient       RESTClient
	rbacClient          RESTClient
	networkingClient    RESTClient
	apiExtensionsClient RESTClient
	pluginsClient       RESTClient
	config              *restclient.Config
}

type crdInfo struct {
	version    string
	group      string
	pluralName string
	namespaced bool
}

// RESTClient is an interface for REST operations used in this file.
type RESTClient interface {
	Delete() *restclient.Request
	Put() *restclient.Request
	Get() *restclient.Request
}

// NewResourceVerber creates a new resource verber that uses the given client for performing operations.
func NewResourceVerber(client, appsClient, batchClient, betaBatchClient, autoscalingClient, storageClient, rbacClient, networkingClient, apiExtensionsClient, pluginsClient RESTClient, config *restclient.Config) clientapi.ResourceVerber {
	return &resourceVerber{client, appsClient,
		batchClient, betaBatchClient, autoscalingClient, storageClient, rbacClient, networkingClient, apiExtensionsClient, pluginsClient, config}
}

func (verber *resourceVerber) Delete(kind string, namespaceSet bool, namespace string, name string) error {
	client, resourceSpec, err := verber.getResourceSpecFromKind(kind, namespaceSet)
	if err != nil {
		return err
	}

	defaultPropagationPolicy := v1.DeletePropagationForeground
	defaultDeleteOptions := &v1.DeleteOptions{
		PropagationPolicy: &defaultDeleteOptions,
	}

	req := client.Delete().Resource()

	if resourceSpec.Namespaced {

	}

	req req.Do(context.TODO()).Error()
}
