package handler

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/kubernetes/dashboard/src/app/backend/resource/deployment"

	"github.com/taidevops/dashboard/src/app/backend/api"
	"github.com/taidevops/dashboard/src/app/backend/errors"
	"github.com/taidevops/dashboard/src/app/backend/handler/parser"
)

type APIHandler struct {
}

func CreateHTTPAPIHandler() (http.Handler, error) {
	apiHandler := APIHandler{}
	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.GET("csrftoken/{action}").
			To(apiHandler.handleGetCsrfToken).
			Writes(api.CsrfToken{}))

	apiV1Ws.Route(
		apiV1Ws.POST("/appdeployment").
			To(apiHandler.handleDeploy).
			Reads(deployment.AppDeploymentSpec{}).
			Writes(deployment.AppDeploymentSpec{}))


	return wsContainer, nil
}

func (apiHandler *APIHandler) handleGetClusterRoleList(request *restful.Request, response *restful.Response) {
	k8sClient, err := apiHandler.cManager.Client(request)
	if err != nil {
		errors.HandleInternalError(response, err)
		return
	}

	dataSelect := parser.ParseDataSelectPathParameter(request)
	result, err := clusterrole.GetClusterRoleList(k8sClient, dataSelect)
	if err != nil {
		errors.HandleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func (apiHandler *APIHandler) handleGetCsrfToken(request *restful.Request, response *restful.Response) {

}
