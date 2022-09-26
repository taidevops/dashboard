package common

import (
	api "k8s.io/api/core/v1"
)

type PodInfo struct {
	Current int32 `json:"current"`

	Desired *int32 `json:"desired,omitempty"`

	Running int32 `json:"running"`

	Pending int32 `json:"pending"`

	Failed int32 `json:"failed"`

	Succeeded int32 `json:"succeeded"`

	Warnings []Event `json:"warnings"`
}