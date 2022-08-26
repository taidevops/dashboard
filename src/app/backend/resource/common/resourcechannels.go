package common

import (
	"context"

	v1 "k8s.io/api/core/v1"
)

type ResourceChannels struct {
	ReplicationControllerList ReplicationControllerListChannel
}

type ServiceListChannel struct {
	List  chan *v1.ServiceList
	Error chan error
}

type IngressListChannel struct {
	List  chan *networkingv1.IngressList
	Error chan error
}

type ReplicationControllerListChannel struct {
	List  chan *v1.ReplicationControllerList
	Error chan error
}
