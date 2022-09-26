package clusterrole

import (

)

type ClusterRoleList struct {
	ListMeta api.ListMeta  `json:"listMeta"`
	Items    []ClusterRole `json:"items"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type ClusterRole struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
}

func GetClusterRoleList(client kubernetes.Interface, dsQuery *dataselect.DataSelectQuery) (*ClusterRoleList, error) {
	channels := &common.ResourceChannels{
		ClusterRoleList: common.
	}
}