package api

import (
	"github.com/emicklei/go-restful/v3"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

const (
	CsrfTokenSecretName = "kubernetes-dashboard-csrf"

	CsrfTokenSecretData = "csrf"
)

type ClientManager interface {
	Client(req *restful.Request) (kubernetes.Interface, error)
}

type ResourceVerber interface {
	Put(kind string, namespaceSet bool, namespace string, name string,
		object *runtime.Unknown) error
}

type CanIResponse struct {
	Allowed bool `json:"allowed"`
}

type CsrfTokenManager interface {
	Token() string
}
