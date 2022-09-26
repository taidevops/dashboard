package integration

import (
	"github.com/taidevops/dashboard/src/app/backend/integration/api"
)

// IntegrationsGetter is responsible for listing all supported integrations.
type IntegrationsGetter interface {
	// List returns list of all supported integrations.
	List() []api.Integration
}

// List implements integration getter interface. See IntegrationsGetter for
// more information.
func (self *integrationManager) List() []api.Integration {
	result := make([]api.Integration, 0)

	// Append all types of integrations
	result = append(result, self.Metric().List()...)

	return result
}
