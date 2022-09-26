package common

import (
	"bytes"

	api "k8s.io/api/core/v1"
)

type Endpoint struct {
	Host string `json:"host"`

	Ports []ServicePort `json:"ports"`
}

//
func GetExternalEndpoints(service *api.Service) []Endpoint {
	externalEndpoints := make([]Endpoint, 0)
	if service.Spec.Type == api.ServiceTypeLoadBalancer {
		for _, ingress := range service.Status.LoadBalancer.Ingress {
			externalEndpoints = append(externalEndpoints, )
		}
	}
}

func getExternalEndpoint(ingress api.LoadBalancerIngress, ports []api.ServicePort) Endpoint {
	var host string
	if ingress.Hostname != "" {
		host = ingress.Hostname
	} else {
		host = ingress.IP
	}
	return Endpoint{
		Host: host,
		Ports: GetServicePorts(ports),
	}
}