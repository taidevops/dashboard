package api

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// CsrfToken is used to secure requests from CSRF attacks
type CsrfToken struct {
	// Token generated on request for validation
	Token string `json:"token"`
}

// ObjectMeta is metadata about an instance of a resource.
type ObjectMeta struct {
	// Name is unique within a namespace. Name is primarily intended for creation
	// idempotence and configuration definition.
	Name string `json:"name,omitempty"`

	// Namespace defines the space within which name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	// Label keys are of the form:
	//     label-key ::= prefixed-name | name
	//     prefixed-name ::= prefix '/' name
	//     prefix ::= DNS_SUBDOMAIN
	//     name ::= DNS_LABEL
	// The prefix is optional.  If the prefix is not specified, the key is assumed to be private
	// to the user.  Other system components that wish to use labels must specify a prefix.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are unstructured key value data stored with a resource that may be set by
	// external tooling. They are not queryable and should be preserved when modifying
	// objects.  Annotation keys have the same formatting restrictions as Label keys. See the
	// comments on Labels for details.
	Annotations map[string]string `json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp v1.Time `json:"creationTimestamp,omitempty"`

	// UID is a type that holds unique ID values, including UUIDs.  Because we
	// don't ONLY use UUIDs, this is an alias to string.  Being a type captures
	// intent and helps make sure that UIDs and names do not get conflated.
	UID types.UID `json:"uid,omitempty"`
}

// TypeMeta describes an individual object in an API response or request with strings representing
// the type of the object.
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// In smalllettercase.
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
	Kind ResourceKind `json:"kind,omitempty"`

	// Scalable represents whether or not an object is scalable.
	Scalable bool `json:"scalable,omitempty"`

	// Restartable represents whether or not an object is restartable.
	Restartable bool `json:"restartable,omitempty"`
}

// ListMeta describes list of objects, i.e holds information about pagination options set for the list.
type ListMeta struct {
	// Total number of items on the list. Used for pagination.
	TotalItems int `json:"totalItems"`
}

// ObjectMeta returns internal endpoint name for the given service properties. e.g.,
// NewObjectMeta creates a new instance of ObjectMeta struct based on K8s object meta.
func NewObjectMeta(k8SObjectMeta metaV1.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name:              k8SObjectMeta.Name,
		Namespace:         k8SObjectMeta.Namespace,
		Labels:            k8SObjectMeta.Labels,
		CreationTimestamp: k8SObjectMeta.CreationTimestamp,
		Annotations:       k8SObjectMeta.Annotations,
		UID:               k8SObjectMeta.UID,
	}
}

// NewTypeMeta creates new type mete for the resource kind.
func NewTypeMeta(kind ResourceKind) TypeMeta {
	return TypeMeta{
		Kind:        kind,
		Scalable:    kind.Scalable(),
		Restartable: kind.Restartable(),
	}
}

// ResourceKind is an unique name for each resource. It can used for API discovery and generic
// code that does things based on the kind. For example, there may be a generic "deleter"
// that based on resource kind, name and namespace deletes it.
type ResourceKind string

const (
	ResourceKindConfigMap    = "configmap"
	ResourceKindDaemonSet    = "daemonset"
	ResourceKindStorageClass = "storageclass"
)

func (k ResourceKind) Scalable() bool {
	return false
}

func (k ResourceKind) Restartable() bool {
	return false
}

// ClientType represents type of client that is used to perform generic operations on resources.
// Different resources belong to different client, i.e. Deployments belongs to extension client
// and StatefulSets to apps client.
type ClientType string

const (
	ClientTypeDefault    = "restclient"
	ClientTypeAppsClient = "appsclient"
)

type APIMapping struct {
	Resource string

	ClientType ClientType

	Namespaced bool
}

// KindToAPIMapping is the mapping from resource kind to K8s apiserver API path. This is mostly pluralization, because
// Kubernetes apiserver uses plural paths and this project singular.
// Must be kept in sync with the list of supported kinds.
// See: https://kubernetes.io/docs/reference/
var KindToAPIMapping = map[string]APIMapping{
	ResourceKindConfigMap: {"configmaps", ClientTypeDefault, true},
	ResourceKindDaemonSet: {"daemonsets", ClientTypeAppsClient, true},
}
