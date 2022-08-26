package common

type ResourceStatus struct {
	// Number of resources that are currently in running state.
	Running int `json:"running"`

	// Number of resources that are currently in pending state.
	Pending int `json:"pending"`
}