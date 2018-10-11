package j7

const (
	LogLevelError   = iota
	LogLevelWarning = iota
	LogLevelInfo    = iota
	LogLevelVerbose = iota
)

type Logger interface {
	Log(level int, message string)
}
