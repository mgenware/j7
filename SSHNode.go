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
	dir        *dirManager
	config     *SSHConfig
	sshConfig  *ssh.ClientConfig
	sshHostKey ssh.PublicKey
	client     *ssh.Client

	Logger Logger
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
		dir:        &dirManager{},
	}, nil
}

func MustNewSSHNode(config *SSHConfig) *SSHNode {
	node, err := NewSSHNode(config)
	if err != nil {
		panic(err)
	}
	return node
}

func (node *SSHNode) SafeRun(cmd string) ([]byte, error) {
	session, err := node.prepareSession()
	output, err := node.runCore(session, cmd)
	if err != nil {
		if len(output) > 0 {
			node.log(LogLevelWarning, string(output))
		}
		node.log(LogLevelWarning, err.Error())
		node.log(LogLevelWarning, "Reconnecting...")

		// set the previous client to nil
		if node.client != nil {
			err = node.client.Close()
			if err != nil {
				node.log(LogLevelWarning, "Error closing previous client: "+err.Error())
			}
			node.client = nil
		}
		session, err = node.prepareSession()
		if err != nil {
			return nil, err
		}

		output, err = node.runCore(session, cmd)
		if err != nil {
			return output, err
		}
	}
	return output, nil
}

func (node *SSHNode) Run(cmd string) []byte {
	output, err := node.SafeRun(cmd)
	if err != nil {
		panic(err)
	}
	return output
}

func (node *SSHNode) runCore(session *ssh.Session, cmd string) ([]byte, error) {
	node.dir.Next(cmd, false)

	lastDir := node.dir.LastDir()
	if lastDir != "" {
		// TODO: better handling of escaping path
		cmd = "cd '" + lastDir + "' && " + cmd
	}

	return session.CombinedOutput(cmd)
}

func (node *SSHNode) prepareSession() (*ssh.Session, error) {
	if node.client == nil {
		cfg := node.config
		addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
		node.log(LogLevelInfo, fmt.Sprintf("SSH: Connecting to %v\n", addr))

		client, err := ssh.Dial("tcp", addr, node.sshConfig)
		if err != nil {
			return nil, err
		}
		node.client = client
	}

	session, err := node.client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (node *SSHNode) log(logLevel int, message string) {
	if node.Logger != nil {
		node.Logger.Log(logLevel, message)
	}
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
