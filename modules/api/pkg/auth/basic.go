package auth

import (
	authApi "github.com/taidevops/dashboard/src/app/backend/auth/api"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Implements Authenticator interface
type basicAuthenticator struct {
	username string
	password string
}

// GetAuthInfo implements Authenticator interface. See Authenticator for more information.
func (self *basicAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	return api.AuthInfo{
		Username: self.username,
		Password: self.password,
	}, nil
}

// NewBasicAuthenticator returns Authenticator based on LoginSpec.
func NewBasicAuthenticator(spec *authApi.LoginSpec) authApi.Authenticator {
	return &basicAuthenticator{
		username: spec.Username,
		password: spec.Password,
	}
}
