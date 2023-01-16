package certificates

import (
	"crypto/tls"
	"log"
	"os"

	"k8s.io/dashboard/certificates/api"
)

// Manager is used to implement cert/api/types.Manager interface. See Manager for more information.
type Manager struct {
	creator      api.Creator
	certDir      string
	autogenerate bool
}

// GetCertificates implements Manager interface. See Manager for more information.
func (self *Manager) GetCertificates() ([]tls.Certificate, error) {
	// Make the autogenerate the top priority option.
	if self.autogenerate {
		key := self.creator.GenerateKey()
		cert := self.creator.GenerateCertificate(key)
		log.Println("Successfully created certificates")
		keyPEM, certPEM, err := self.creator.KeyCertPEMBytes(key, cert)
		if err != nil {
			return []tls.Certificate{}, err
		}
		certificate, err := tls.X509KeyPair(certPEM, keyPEM)
		return []tls.Certificate{certificate}, err
	}

	// When autogenerate is disabled and provided cert files exist use them.
	if self.keyFileExists() && self.certFileExists() && !self.autogenerate {
		log.Println("Certificates already exist. Returning.")
		certificate, err := tls.LoadX509KeyPair(
			self.path(self.creator.GetCertFileName()),
			self.path(self.creator.GetKeyFileName()),
		)

		return []tls.Certificate{certificate}, err
	}

	return nil, nil
}