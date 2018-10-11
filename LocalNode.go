package j7

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
	dir := node.dir.Next(cmd, true)
	if dir != "" {
		// Unlike SSH session, once we set the working dir to a value, we don't need to reset it on subsequent commands.
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
