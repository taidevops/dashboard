package node

import (
	"context"

	v1 "k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"

	"github.com/kubernetes/dashboard/src/app/backend/api"
)

type NodeList struct {
	ListMeta api.ListMeta `json:"listMeta"`
}

type Node struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
}

