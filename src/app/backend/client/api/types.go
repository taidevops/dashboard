package api

import (
	"github.com/emicklei/go-restful/v3"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

const (
	CsrfTokenSecretName = "kubernetes-dashboard-csrf"

	CsrfTokenSecretData = "csrf"
)

// ClientManager is responsible for initializing and creating clients to communicate with
// kubernetes apiserver on demand.
type ClientManager interface {
	Client(req *restful.Request) (kubernetes.Interface, error)
	InsecureClient() kubernetes.Interface
	APIExtensionsClient(req *restful.Request) (apiextensionsclientset.Interface, error)
}

// ResourceVerber is responsible for performing generic CRUD operations on all supported resources.
type ResourceVerber interface {
	Put(kind string, namespaceSet bool, namespace string, name string,
		object *runtime.Unknown) error
	Get(kind string, namespaceSet bool, namespace string, name string) (runtime.Object, error)
	Delete(kind string, namespaceSet bool, namespace string, name string) error
}

// CanIResponse is used to as response to check whether or not user is allowed to access given endpoint.
type CanIResponse struct {
	Allowed bool `json:"allowed"`
}

// CsrfTokenManager is responsible for generating, reading and updating token stored in a secret.
type CsrfTokenManager interface {
	// Token returns current csrf token used for csrf signing.
	Token() string
}
