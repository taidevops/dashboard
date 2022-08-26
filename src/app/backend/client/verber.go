package client

import (

)

type resourceVerber struct {
	client RESTClient
	appsClient RESTClient
}

// RESTClient is an interface for REST operations used in this file.
type RESTClient interface {
	Delete() *restclient.Request
	Put() *restclient.Request
	Get() *restclient.Request
}
