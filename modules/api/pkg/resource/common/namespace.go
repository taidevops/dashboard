package common

import api "k8s.io/api/core/v1"

type NamespaceQuery struct {
	namespaces []string
}

func NewSameNamespaceQuery(namespace string) *NamespaceQuery {
	return &NamespaceQuery{[]string{namespace}}
}

func NewNamespaceQuery(namespaces []string) *NamespaceQuery {
	return &NamespaceQuery{namespaces}
}

