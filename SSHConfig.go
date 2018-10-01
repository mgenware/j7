package lun

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

type SSHConfig struct {
	Host string
	User string
	Port int
	Auth []ssh.AuthMethod
}

func NewKeyBasedAuth(keyFile string) ([]ssh.AuthMethod, error) {
	keyBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return []ssh.AuthMethod{ssh.PublicKeys(signer)}, nil
}

func MustNewKeyBasedAuth(keyFile string) []ssh.AuthMethod {
	auth, err := NewKeyBasedAuth(keyFile)
	if err != nil {
		panic(err)
	}
	return auth
}
