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

func (node *LocalNode) RunOrError(cmd string) ([]byte, error) {
	// Get the next working dir to be set
	wd := node.dir.NextWD(cmd, true)
	if wd != "" {
		// Unlike SSH session, once we set the working dir to a value, we don't need to set it on subsequent commands.
		err := os.Chdir(wd)
		if err != nil {
			return []byte("Cannot change working dir to: \"" + wd + "\""), err
		}
	}

	output, err := node.execCore("bash", "-c", cmd)
	if err != nil {
		return output, err
	}
	return output, nil
}

func (node *LocalNode) execCore(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}
