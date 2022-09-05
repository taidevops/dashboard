package storageclass

import (
	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/taidevops/dashboard/src/app/backend/api"
	"github.com/taidevops/dashboard/src/app/backend/resource/dataselect"
	"k8s.io/client-go/kubernetes"
)

type StorageClassList struct {
	ListMeta api.ListMeta
}

type StorageClass struct {
	ObjectMeta  api.ObjectMeta    `json:"objectMeta"`
	TypeMeta    api.TypeMeta      `json:"typeMeta"`
	Provisioner string            `json:"provisiner"`
	Parameters  map[string]string `json:"parameters"`
}

func GetStorageClassList(client kubernetes.Interface, dsQuery *dataselect.DataSelectQuery) (
	*StorageClassList, error) {
	channels := &common.ResourceChannels{}
}
