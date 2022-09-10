package integration

import (
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