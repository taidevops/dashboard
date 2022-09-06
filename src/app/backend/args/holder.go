package args

import (
	"net"
)

var Holder = &holder{}

type holder struct {
	insecurePort            int
	port                    int
	tokenTTL                int
	metricClientCheckPeriod int

	insecureBindAddress net.IP
	bindAddress         net.IP

	defaultCertDir       string
	certFile             string
	keyFile              string
	apiServerHost        string
	metricsProvider      string
	heapsterHost         string
	sidecarHost          string
	kubeConfigFile       string
	systemBanner         string
	systemBannerSeverity string
	apiLogLevel          string
	namespace            string

	authenticationMode []string

	autoGenerateCertificates  bool
	enableInsecureLogin       bool
	disableSettingsAuthorizer bool

	enableSkipLogin bool

	localeConfig string
}

func (self *holder) GetInsecurePort() int {
	return self.insecurePort
}

func (self *holder) GetPort() int {
	return self.port
}

func (self *holder) GetTokenTTL() int {
	return self.tokenTTL
}

func (self *holder) GetMetricClientCheckPeriod() int {
	return self.metricClientCheckPeriod
}

func (self *holder) GetInsecureBindAddress() net.IP {
	return self.insecureBindAddress
}

// GetEnableInsecureLogin 'enable-insecure-login' argument of Dashboard binary.
func (self *holder) GetEnableInsecureLogin() bool {
	return self.enableInsecureLogin
}

// GetEnableSkipLogin 'enable-skip-login' argument of Dashboard binary.
func (self *holder) GetEnableSkipLogin() bool {
	return self.enableSkipLogin
}
