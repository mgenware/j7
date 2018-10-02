package lun

type Node interface {
	SafeExec(cmd string) ([]byte, error)
	Exec(cmd string) []byte
}
