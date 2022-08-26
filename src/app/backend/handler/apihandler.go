package handler

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"github.com/kubernetes/dashboard/src/app/backend/errors"
	"github.com/kubernetes/dashboard/src/app/backend/handler/parser"
)

type APIHandler struct {
}

func CreateHTTPAPIHandler() (http.Handler, error) {

	wsContainer := restful.NewContainer()
	return wsContainer, nil
}

func (apiHandler *APIHandler) handleGetClusterRoleList(request *restful.Request, response *restful.Response) {
	k8sClient, err := apiHandler.cManager.Client(request)
	if err != nil {
		errors.HandleInternalError(response, err)
		return
	}

	dataSet := parser.ParseDataSelectPathParameter(request)
	result, err := 
}