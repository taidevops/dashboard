package jwe

import (
	"crypto/rsa"

	jose "gopkg.in/square/go-jose.v2"
)

const (
	holderMapKeyEntry  = "priv"
	holderMapCertEntry = "pub"
)

type KeyHolder interface {
	Encrypter() jose.Encrypter
	Key() *rsa.PrivateKey
	Refresh()
}

type rsaKeyHolder struct {
	key *rsa.PrivateKey
	synchroizer asyn
}
