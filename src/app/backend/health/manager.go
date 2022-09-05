package health

import (
	client "github.com/taidevops/dashboard/src/app/backend/client/api"
	health "github.com/taidevops/dashboard/src/app/backend/health/api"
)

// HealthManager is a structure containing all system banner manager members.
type HealthManager struct {
	client client.ClientManager
}

// NewHealthManager creates new settings manager.
func NewHealthManager(client client.ClientManager) HealthManager {
	return HealthManager{
		client: client,
	}
}

// Get implements HealthManager interface. Check it for more information.
func (sbm *HealthManager) Get() health.Health {
	_, err := sbm.client.InsecureClient().Discovery().ServerVersion()
	return health.Health{
		Running: err == nil,
	}
}
