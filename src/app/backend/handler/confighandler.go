package handler

import "net/http"

// AppHandler is an application handler.
type AppHandler func(http.ResponseWriter, *http.Request) (int, error)
