package handler

import (
	"io"

	restful "github.com/emicklei/go-restful/v3"

	"github.com/taidevops/dashboard/src/app/backend/errors"
)

func handleDownload(response *restful.Response, result io.ReadCloser) {
	response.AddHeader(restful.HEADER_ContentType, "text/plain")
	defer result.Close()
	_, err := io.Copy(response, result)
	if err != nil {
		errors.HandleInternalError(response, err)
	}
}
