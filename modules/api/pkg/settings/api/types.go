package api

import (
	"encoding/json"

	"k8s.io/client-go/kubernetes"
)

const (
	SettingsConfigMapName = "kubernetes-dashboard-settings"
)

type SettingsManager interface {
	// GetGlobalSettings gets current global settings from config map.
	GetGlobalSettings(client kubernetes.Interface) (s Settings)
}

type Settings struct {
	ClusterName string `json:"clusterName"`
	ItemsPerPage int `json:"itemsPerPage"`
	
}