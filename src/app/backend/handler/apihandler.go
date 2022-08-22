package handler

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type APIHandler struct {
}

func CreateHTTPAPIHandler() (http.Handler, error) {
	wsContainer := restful.NewContainer()
	return wsContainer, nil
}
