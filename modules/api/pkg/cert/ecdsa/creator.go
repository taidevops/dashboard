package ecdsa

import (
	"crypto/elliptic"

	certapi "github.com/taidevops/dashboard/src/app/backend/cert/api"
)

// Implements certificate Creator interface. See Creator for more information.
type ecdsaCreator struct {
	keyFile  string
	certFile string
	curve    elliptic.Curve
}

func (self *ecdsaCreator) init() {
	if len(self.certFile) == 0 {
		self.certFile = certapi.DashboardCertName
	}

	if len(self.keyFile) == 0 {
		self.keyFile = certapi.DashboardKeyName
	}
}

// NewECDSACreator creates ECDSACreator instance.
func NewECDSACreator(keyFile, certFile string, curve elliptic.Curve) certapi.Creator {
	creator := &ecdsaCreator{
		curve:    curve,
		keyFile:  keyFile,
		certFile: certFile,
	}

	creator.init()
	return creator
}
