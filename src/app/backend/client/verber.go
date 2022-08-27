package client

import (
	"github.com/taidevops/dashboard/src/app/backend/api"
	restclient "k8s.io/client-go/rest"
)

type resourceVerber struct {
	client     RESTClient
	appsClient RESTClient
}

// RESTClient is an interface for REST operations used in this file.
type RESTClient interface {
	Delete() *restclient.Request
	Put() *restclient.Request
	Get() *restclient.Request
}

type crdInfo struct {
	version    string
	group      string
	pluralName string
	namespaced bool
}

func (verber *resourceVerber) getRESTClientByType(clientType api.ClientType) RESTClient {
	switch clientType {
	case api.ClientTypeAppsClient:
		return verber.appsClient
	default:
		return verber.client
	}
}
