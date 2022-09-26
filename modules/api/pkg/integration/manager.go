package integration

import (
	clientapi "github.com/taidevops/dashboard/src/app/backend/client/api"
	"github.com/taidevops/dashboard/src/app/backend/integration/metric"
)

type IntegrationManager interface {

}

// Implements IntegrationManager interface
type integrationManager struct {
	metric metric.MetricManager
}

func (self *integrationManager) Metric() metric.MetricManager {
	return self.metric
}

func NewIntegrationManager(manager clientapi.ClientManager) IntegrationManager {
	return &integrationManager{
		metric: metric.NewMetricManager(manager),
	}
}