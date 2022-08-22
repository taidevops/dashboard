package health

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct {

}

// Install creates new endpoints for system banner mangement.
func (self *HealthHandler) Install(w http.ResponseWriter, _ *http.Request) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	return http.StatusOK, json.NewEncoder(w).Encode(self.manager.Get())
}

// NewHealthHandler creates HealthManager.
func NewHealthHandler(manager HealthHandler) HealthHandler {
	return HealthHandler{manager: manager}
}