package runners

import "github.com/mgenware/lun"

type Runner interface {
	Run(node lun.Node, cmd string)
	SafeRun(node lun.Node, cmd string)
}
