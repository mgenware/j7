package lun

type Node interface {
	SafeRun(cmd string) ([]byte, error)
	Run(cmd string) []byte
}
