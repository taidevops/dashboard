package integration

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/taidevops/dashboard/src/app/backend/integration/api"
)

// IntegrationHandler manages all endpoints related to integrated applications, such as state.
type IntegrationHandler struct {
	manager IntegrationManager
}

func (self IntegrationHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.GET("/integration/{name}/state").
			To(self.handleGetState).
			Writes(api.IntegrationState{}))
}

func (self IntegrationHandler) handleGetState(request *restful.Request, response *restful.Response) {

}
