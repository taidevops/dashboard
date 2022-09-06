package storageclass

import (
	"github.com/taidevops/dashboard/src/app/backend/resource/common"
	"github.com/taidevops/dashboard/src/app/backend/api"
	"github.com/taidevops/dashboard/src/app/backend/resource/dataselect"
	storage "k8s.io/api/storage/v1"
	"k8s.io/client-go/kubernetes"
)

type StorageClassList struct {
	ListMeta api.ListMeta `json:"listMeta"`
	Items []StorageClass `json:"items"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type StorageClass struct {
	ObjectMeta  api.ObjectMeta    `json:"objectMeta"`
	TypeMeta    api.TypeMeta      `json:"typeMeta"`
	Provisioner string            `json:"provisiner"`
	Parameters  map[string]string `json:"parameters"`
}

func GetStorageClassList(client kubernetes.Interface, dsQuery *dataselect.DataSelectQuery) (
	*StorageClassList, error) {
	channels := &common.ResourceChannels{
		StorageClassList: common.Get,
	}
}

func toStorageClass(storageClass *storage.StorageClass) StorageClass {
	return StorageClass{
		ObjectMeta: api.NewObjectMeta(storageClass.ObjectMeta),
		TypeMeta: api.NewTypeMeta(api.ResourceKindStorageClass),
		Provisioner: storageClass.Provisioner,
		Parameters: storageClass.Parameters,
	}
}