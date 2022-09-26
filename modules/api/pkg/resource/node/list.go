package node

import (
	"context"

	v1 "k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"

	"github.com/kubernetes/dashboard/src/app/backend/api"
	metricapi "github.com/taidevops/dashboard/src/app/backend/integration/metric/api"
)

type NodeList struct {
	ListMeta api.ListMeta `json:"listMeta"`
	Nodes []Node `json:"nodes"`
	CumulativeMetrics []metricapi.Me
}

type Node struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
}

