package loggers

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mgenware/lun"
)

type ConsoleLogger struct {
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (c *ConsoleLogger) Log(level int, message string) {
	if level == lun.LogLevelVerbose {
		fmt.Println(message)
	} else {
		var console func(format string, a ...interface{})
		switch level {
		case lun.LogLevelError:
			console = color.Red
		case lun.LogLevelWarning:
			console = color.Yellow
		default:
			console = color.Cyan
		}

		console(message)
	}
}
