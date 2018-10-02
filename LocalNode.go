package lun

import (
	"os/exec"
)

type localNode struct {
}

func (node *localNode) SafeExec(cmd string) ([]byte, error) {
	output, err := node.execCore("bash", "-c", cmd)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (node *localNode) Exec(cmd string) []byte {
	output, err := node.SafeExec(cmd)
	if err != nil {
		panic(err)
	}
	return output
}

func (node *localNode) execCore(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}

// LocalNode is a static instance of internal localNode.
var LocalNode = &localNode{}
