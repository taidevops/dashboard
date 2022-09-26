package common

import (
	"context"
	apps "k8s.io/api/apps/v1"

	v1 "k8s.io/api/core/v1"

	client "k8s.io/client-go/kubernetes"
)

type ResourceChannels struct {
	ReplicationControllerList ReplicationControllerListChannel

	ReplicaSetList Re
}

type ServiceListChannel struct {
	List  chan *v1.ServiceList
	Error chan error
}

func GetServiceListChannel(client client.Interface, nsQuery *NamespaceQuery,
	numReads int) ServiceListChannel {

	channel := ServiceListChannel{
		List: make(chan *v1.ServiceList, numReads),
		Error: make(chan error, numReads),
	}
	go func() {

	}()
}

type IngressListChannel struct {
	List  chan *networkingv1.IngressList
	Error chan error
}

type ReplicationControllerListChannel struct {
	List  chan *v1.ReplicationControllerList
	Error chan error
}

type ReplicaSetListChannel struct {
	List chan *apps.ReplicaSetList
}
