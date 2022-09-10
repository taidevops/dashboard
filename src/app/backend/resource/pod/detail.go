package pod

import (
	"github.com/taidevops/dashboard/src/app/backend/api"
)

type PodDetail struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
	TypeMeta   api.TypeMeta   `json:"typeMeta"`
	PodPhase   string         `json:"podPhase"`
}

type Container struct {
	Name string `json:"name"`

	Image string `json:"image"`
}
