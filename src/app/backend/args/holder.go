package args

import (
	"net"
)

var Holder = &holder{}

type holder struct {
	insecurePort int
	port         int
	tokenTTL     int

	insecureBindAddress net.IP
	bindAddress         net.IP

	enableInsecureLogin bool

	enableSkipLogin bool
}

// GetEnableInsecureLogin 'enable-insecure-login' argument of Dashboard binary.
func (self *holder) GetEnableInsecureLogin() bool {
	return self.enableInsecureLogin
}

// GetEnableSkipLogin 'enable-skip-login' argument of Dashboard binary.
func (self *holder) GetEnableSkipLogin() bool {
	return self.enableSkipLogin
}
