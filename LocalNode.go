package lun

import (
	"os"
	"os/exec"
)

// LocalNode is used for running commands locally.
type LocalNode struct {
	dir *dirManager
}

func NewLocalNode() *LocalNode {
	return &LocalNode{
		dir: &dirManager{},
	}
}

func (node *LocalNode) SafeRun(cmd string) ([]byte, error) {
	lastDir := node.dir.LastDir()
	if lastDir != "" {
		output, err := node.execCore("cd", lastDir)
		if err != nil {
			return output, err
		}
	}

	dir := node.dir.Next(cmd, true)
	if dir != "" {
		err := os.Chdir(dir)
		if err != nil {
			return nil, err
		}
	}

	output, err := node.execCore("bash", "-c", cmd)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (node *LocalNode) Run(cmd string) []byte {
	output, err := node.SafeRun(cmd)
	if err != nil {
		panic(err)
	}
	return output
}

func (node *LocalNode) execCore(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}
