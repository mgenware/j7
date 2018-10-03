package runners

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mgenware/lun"
)

type ConsoleRunner struct {
}

func NewConsoleRunner() *ConsoleRunner {
	return &ConsoleRunner{}
}

func (r *ConsoleRunner) Run(node lun.Node, cmd string) {
	r.run(false, node, cmd)
}

func (r *ConsoleRunner) SafeRun(node lun.Node, cmd string) error {
	return r.run(true, node, cmd)
}

func (r *ConsoleRunner) run(ignore bool, node lun.Node, cmd string) error {
	if ignore {
		color.Cyan("🚙 " + cmd)
	} else {
		color.Yellow("🚗 " + cmd)
	}
	output, err := node.SafeExec(cmd)
	if err != nil {
		if len(output) > 0 {
			color.Red(string(output))
		}
		color.Red(err.Error())
		if !ignore {
			panic(err)
		}
	} else {
		if len(output) > 0 {
			fmt.Print(string(output))
		}
	}
	return err
}
