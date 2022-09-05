package metric

import (
	"fmt"
	"log"
	"time"

	clientapi "github.com/taidevops/dashboard/src/app/backend/client/api"
	integrationapi "github.com/taidevops/dashboard/src/app/backend/integration/api"
	metricapi "github.com/taidevops/dashboard/src/app/backend/integration/metric/api"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/taidevops/dashboard/src/app/backend/integration/metric/sidecar"
)

type MetricManager interface {
	AddClient(metricapi.MetricClient) MetricManager

	Client() metricapi.MetricClient

	Enable(integrationapi.IntegrationID) error

	EnableWithRetry(id integrationapi.IntegrationID, period time.Duration)

	List() []integrationapi.Integration

	ConfiguresSidecar(host string) MetricManager
}

type metricManager struct {
	manager clientapi.ClientManager
	clients map[integrationapi.IntegrationID]metricapi.MetricClient
	active metricapi.MetricClient
}

func (self *metricManager) AddClient(client metricapi.MetricClient) MetricManager {
	if client != nil {
		self.clients[client.ID()] = client
	}

	return self
}

func (self *metricManager) Client() metricapi.MetricClient {
	return self.active
}

func (self *metricManager) Enable(id integrationapi.IntegrationID) error {
	metricClient, exists := self.clients[id]
	if !exists {
		return fmt.Errorf("No metric client found for integration id: %s", id)
	}

	err := metricClient.HealthCheck()
	if err != nil {
		return fmt.Errorf("Health check failed: %s", err.Error())
	}

	self.active = metricClient
	return nil
}

func (self *metricManager) EnableWithRetry(id integrationapi.IntegrationID, period time.Duration) {
	go wait.Forever(func() {
		metricClient, exists := self.clients[id]
		if !exists {
			log.Printf("Metric client with given id %s does not exist.", id)
			return
		}

		err := metricClient.HealthCheck()
		if err != nil {
			self.active = nil
			log.Printf("Metric client health check failed: %s. Retrying in %d seconds.", err, period)
		}

	}, period*time.Second)
}

func (self *metricManager) List() []integrationapi.Integration {
	result := make([]integrationapi.Integration, 0)
	for _, c := range self.clients {
		result = append(result, c.(integrationapi.Integration))
	}

	return result
}

func (self *metricManager) ConfigureSidecar(host string) MetricManager {
	kubeClient := self.manager.InsecureClient()
	metricClient, err := sidecar
}