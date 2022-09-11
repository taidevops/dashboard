package auth

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"

	authApi "github.com/taidevops/dashboard/src/app/backend/auth/api"
	"github.com/taidevops/dashboard/src/app/backend/errors"
	"github.com/taidevops/dashboard/src/app/backend/validation"
)

type AuthHandler struct {
	manager authApi.AuthManager
}

// Install creates new endpoints for dashboard auth, such as login. It allows user to log in to dashboard using
// one of the supported methods. See AuthManager and Authenticator for more information.
func (self AuthHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/login").
			To(self.handleLogin).
			Reads(authApi.LoginSpec{}).
			Writes(authApi.AuthResponse{}))
	ws.Route(
		ws.GET("/login/status").
			To(self.handleLoginStatus).
			Writes(validation.LoginStatus{}))
	ws.Route(
		ws.POST("/token/refresh").
			Reads(authApi.TokenRefreshSpec{}).
			To(self.handleJWETokenRefresh).
			Writes(authApi.AuthResponse{}))
	ws.Route(
		ws.GET("/login/modes").
			To(self.handleLoginModes).
			Writes(authApi.LoginModesResponse{}))
	ws.Route(
		ws.GET("/login/skippable").
			To(self.handleLoginSkippable).
			Writes(authApi.LoginSkippableResponse{}))
}

func (self AuthHandler) handleLogin(request *restful.Request, response *restful.Response) {
	loginSpec := new(authApi.LoginSpec)
	if err := request.ReadEntity(loginSpec); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	loginResponse, err := self.manager.Login(loginSpec)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, loginResponse)
}

func (self *AuthHandler) handleLoginStatus(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, validation.ValidateLoginStatus(request))
}

func (self *AuthHandler) handleJWETokenRefresh(request *restful.Request, response *restful.Response) {
	tokenRefreshSpec := new(authApi.TokenRefreshSpec)
	if err := request.ReadEntity(tokenRefreshSpec); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	refreshedJWEToken, err := self.manager.Refresh(tokenRefreshSpec.JWEToken)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(errors.HandleHTTPError(err), err.Error()+"\n")
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, &authApi.AuthResponse{
		JWEToken: refreshedJWEToken,
		Errors:   make([]error, 0),
	})
}

func (self *AuthHandler) handleLoginModes(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, authApi.LoginModesResponse{Modes: self.manager.AuthenticationModes()})
}

func (self *AuthHandler) handleLoginSkippable(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, authApi.LoginSkippableResponse{Skippable: self.manager.AuthenticationSkippable()})
}

// NewAuthHandler created AuthHandler instance.
func NewAuthHandler(manager authApi.AuthManager) AuthHandler {
	return AuthHandler{manager: manager}
}
