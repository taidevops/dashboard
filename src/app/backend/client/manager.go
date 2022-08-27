package client

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/taidevops/dashboard/src/app/backend/args"
	clientapi "github.com/taidevops/dashboard/src/app/backend/client/api"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Dashboard UI default values for client configs.
const (
	DefaultQPS = 1e6

	// Header name that contains token used for authorization. See TokenManager for more information.
	JWETokenHeader = "jweToken"
)

var Version = "UNKNOWN"

type clientManager struct {
	csrKey string
}

func (self *clientManager) Client(req *restful.Request) (kubernetes.Interface, error) {
	if req == nil {
		return nil, errors.NewBadRequest("request can not be nil")
	}

	if self.isSecureModeEnabled(req) {
		return self.secureClient(req)
	}

	return self.InsecureClient(), nil
}

// APIExtensionsClient returns an API Extensions client. In case dashboard login is enabled and
// option to skip login page is disabled only secure client will be returned, otherwise insecure
// client will be used.
func (self *clientManager) APIExtensionsClient(req *restful.Request) (apiextensionsclientset.Interface, error) {
	if req == nil {
		return nil, errors.NewBadRequest("request can not be nil!")
	}

	if self.isSecureModeEnabled(req) {
		return self.secureAPIExtensionsClient(req)
	}

	return self.InsecureAPIExtensionsClient(), nil
}

// Config returns a rest config. In case dashboard login is enabled and option to skip
// login page is disabled only secure config will be returned, otherwise insecure config will be
// used.
func (self *clientManager) Config(req *restful.Request) (*rest.Config, error) {
	if req == nil {
		return nil, errors.NewBadRequest("request can not be nil")
	}

	if self.isSecureModeEnabled(req) {
		return self.secureConfig(req)
	}

	return self.InsecureConfig(), nil
}

// InsecureClient returns kubernetes client that was created without providing auth info. It uses
// permissions granted to service account used by dashboard or kubeconfig file if it was passed
// during dashboard init.
func (self *clientManager) InsecureClient() kubernetes.Interface {
	return self.insecureClient
}

func (self *clientManager) InsecureAPIExtensionsClient() apiextensionsclientset.Interface {
	return self.insecureAPI
}

func (self *clientManager) secureClient(req *restful.Request) (kubernetes.Interface, error) {
	cfg, err := self.secureConfig(req)
}

func (self *clientManager) secureConfig(req *restful.Request) (*rest.Config, error) {
	cmdConfig, err := self.
}

// Checks if request headers contain any auth information without parsing.
func (self *clientManager) containsAuthInfo(req *restful.Request) bool {
	authHeader := req.HeaderParameter("Authorization")
	jweToken := req.HeaderParameter(JWETokenHeader)

	return len(authHeader) > 0 || len(jweToken) > 0
}

func (self *clientManager) isLoginEnabled(req *restful.Request) bool {
	return req.Request.TLS != nil || args.Holder.GetEnableInsecureLogin()
}

// Secure mode means that every request to Dashboard has to be authenticated and privileges
// of Dashboard SA can not be used.
func (self *clientManager) isSecureModeEnabled(req *restful.Request) bool {
	if self.isLoginEnabled(req) && !args.Holder.GetEnableSkipLogin() {
		return true
	}

	return self.isLoginEnabled(req) && args.Holder.GetEnableSkipLogin() && self.containsAuthInfo(req)
}

func NewClientManager(kubeConfigPath, apiserverHost string) clientapi.ClientManager {
	result := &clientManager{}

	result.init()
	return result
}
