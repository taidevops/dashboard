package handler

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"github.com/taidevops/dashboard/src/app/backend/api"
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

	return wsContainer, nil
}

func (apiHandler *APIHandler) handleGetCsrfToken(request *restful.Request, response *restful.Response) {

}
