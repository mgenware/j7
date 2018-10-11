package loggers

import (
	"fmt"

	"github.com/fatih/color"
)

type ConsoleLogger struct {
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (c *ConsoleLogger) Log(level int, message string) {
	if level == j7.LogLevelVerbose {
		fmt.Println(message)
	} else {
		var console func(format string, a ...interface{})
		switch level {
		case j7.LogLevelError:
			console = color.Red
		case j7.LogLevelWarning:
			console = color.Yellow
		default:
			console = color.Cyan
		}

		console(message)
	}
}
