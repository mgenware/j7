package lun

type Runner interface {
	Run(node Node, cmd string)
	SafeRun(node Node, cmd string)
}
