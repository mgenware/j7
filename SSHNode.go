package lun

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

type SSHNode struct {
	config    *SSHConfig
	sshConfig *ssh.ClientConfig

	sshHostKey ssh.PublicKey
}

func NewSSHNode(config *SSHConfig) (*SSHNode, error) {
	if config.Host == "" {
		return nil, fmt.Errorf("config.Host cannot be empty")
	}
	if config.User == "" {
		return nil, fmt.Errorf("config.User cannot be empty")
	}
	if len(config.Auth) == 0 {
		return nil, fmt.Errorf("config.Auth cannot be empty")
	}
	if config.Port == 0 {
		config.Port = 22
	}

	hostKey, err := getHostKey(config.Host)
	if err != nil {
		return nil, err
	}
	sshConfig := &ssh.ClientConfig{
		User:            config.User,
		Auth:            config.Auth,
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	return &SSHNode{
		config:     config,
		sshConfig:  sshConfig,
		sshHostKey: hostKey,
	}, nil
}

func MustNewSSHNode(config *SSHConfig) *SSHNode {
	node, err := NewSSHNode(config)
	if err != nil {
		panic(err)
	}
	return node
}

func (node *SSHNode) Exec(cmd string) ([]byte, error) {
	session, err := node.prepare()
	if err != nil {
		return nil, err
	}

	return session.CombinedOutput(cmd)
}

func (node *SSHNode) prepare() (*ssh.Session, error) {
	cfg := node.config
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), node.sshConfig)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// https://stackoverflow.com/questions/45441735/ssh-handshake-complains-about-missing-host-key
func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		return nil, fmt.Errorf("no hostkey for %s", host)
	}
	return hostKey, nil
}
