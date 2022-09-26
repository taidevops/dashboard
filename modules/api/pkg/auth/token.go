package auth

import (
	authApi "github.com/taidevops/dashboard/src/app/backend/auth/api"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Implements Authenticator interface
type tokenAuthenticator struct {
	token string
}

// GetAuthInfo implements Authenticator interface. See Authenticator for more information.
func (self tokenAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	return api.AuthInfo{
		Token: self.token,
	}, nil
}

// NewTokenAuthenticator returns Authenticator based on LoginSpec.
func NewTokenAuthenticator(spec *authApi.LoginSpec) authApi.Authenticator {
	return &tokenAuthenticator{
		token: spec.Token,
	}
}
