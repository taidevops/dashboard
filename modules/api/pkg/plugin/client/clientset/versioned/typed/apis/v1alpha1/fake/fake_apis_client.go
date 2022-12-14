// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/taidevops/dashboard/src/app/backend/plugin/client/clientset/versioned/typed/apis/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDashboardV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDashboardV1alpha1) Plugins(namespace string) v1alpha1.PluginInterface {
	return &FakePlugins{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDashboardV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
