package node

import (
	"context"
)

type NodeAllocatedResources struct {
	CPURequests int64 `json:"cpuRequests"`
}

type NodeDetail struct {
	Node `json:",inline"`
}

