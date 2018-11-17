package j7

type Node interface {
	Run(cmd string) ([]byte, error)
}
