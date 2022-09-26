package ecdsa

import (
	"crypto/ecdsa"
	"testing"
)

func TestNewECDSACreator(t *testing.T) {
	keyFile := "cert.key"
	certFile := "cert.crt"
	creator := ecdsa.
}