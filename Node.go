package j7

type Node interface {
	RunOrError(cmd string) ([]byte, error)
}
