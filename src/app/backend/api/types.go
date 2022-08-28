package api

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// CsrfToken is used to secure requests from CSRF attacks
type CsrfToken struct {
	// Token generated on request for validation
	Token string `json:"token"`
}

type ObjectMeta struct {
	Name string `json:"name,omitempty"`

	Namespace string `json:"namspace,omitempty"`

	Labels map[string]string `json:"labels"`

	Annotations map[string]string `json:"annotations,omitempty"`

	CreationTimestamp v1.Time `json:"annotation,omitempty"`

	UID types.UID `json:"uid,omitempty"`
}

type ResourceKind string
