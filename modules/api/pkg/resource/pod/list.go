package pod

import "github.com/taidevops/dashboard/src/app/backend/api"

type PodList struct {
	ListMeta api.ListMeta `json:"listMeta"`

	Pods []Pod `json:"pods"`
}

type Pod struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
	TypeMeta   api.TypeMeta   `json:"typeMeta"`

	Status string `json:"status"`
}
