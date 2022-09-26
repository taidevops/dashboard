package health

import (
	"encoding/json"
	"net/http"
)

// HealthHandler manages all endpoints related to system banner management.
type HealthHandler struct {
	manager HealthManager
}

// Install creates new endpoints for system banner management.
func (self *HealthHandler) Install(w http.ResponseWriter, _ *http.Request) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	return http.StatusOK, json.NewEncoder(w).Encode(self.manager.Get())
}

// NewHealthHandler creates HealthHandler.
func NewHealthHandler(manager HealthManager) HealthHandler {
	return HealthHandler{manager: manager}
}
