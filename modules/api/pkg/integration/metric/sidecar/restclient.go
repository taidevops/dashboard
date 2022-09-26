package sidecar

import (
	"context"

	"k8s.io/client-go/rest"
)

type SidecarRESTClient interface {
	Get(path string) RequestInterface
	HealthCheck() error
}

type RequestInterface interface {
	DoRaw(context.Context) ([]byte, error)
	AbsPath(segments ...string) *rest.Request
}
