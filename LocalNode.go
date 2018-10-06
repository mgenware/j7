package lun

import (
	"os"
	"os/exec"
	"strings"

	"github.com/mgenware/lun/lib"

	"github.com/mgenware/go-packagex/stringsx"
)

// LocalNode is used for running commands locally.
type LocalNode struct {
	lastDir string
}

func NewLocalNode() *LocalNode {
	return &LocalNode{}
}

func (node *LocalNode) SafeRun(cmd string) ([]byte, error) {
	if node.lastDir != "" {
		output, err := node.execCore("cd", node.lastDir)
		if err != nil {
			return output, err
		}
	}

	if strings.HasPrefix(cmd, "cd") {
		var dir string
		if len(cmd) == 2 {
			dir = os.Getenv("HOME")
		} else if cmd[2] == ' ' && len(cmd) > 3 {
			dir = strings.TrimSpace(stringsx.SubStringFromStart(cmd, 3))
		}

		if dir != "" {
			dir = lib.EvaluatePath(dir)
			err := os.Chdir(dir)
			if err != nil {
				return nil, err
			}
			node.lastDir = dir
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
