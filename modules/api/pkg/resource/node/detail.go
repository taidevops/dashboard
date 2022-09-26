package node

type NodeAllocatedResources struct {
	CPURequests int64 `json:"cpuRequests"`

	CPURequestsFraction float64 `json:"cpuRequestsFraction"`

	CPULimits int64 `json:"cpuLimits"`
}

type NodeDetail struct {
	Node `json:",inline"`
}
