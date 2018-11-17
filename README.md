# j7

[![Build Status](https://travis-ci.org/mgenware/j7.svg?branch=master)](http://travis-ci.org/mgenware/j7)

Shell scripting in Go

## Installation
```sh
go get github.com/mgenware/j7
```

## Examples
Checking if a command is installed and performing an install command if necessary. (assuming macOS with homebrew installed)
```go
package main

import (
	"fmt"
	"os/exec"

	"github.com/mgenware/j7"
	"github.com/mgenware/j7/loggers"
)

func main() {
	t := j7.NewTunnel(j7.NewLocalNode(), loggers.NewConsoleLogger())

	_, err := exec.LookPath("tree")
	if err != nil {
		t.Logger().Log(j7.LogLevelError, "tree is not installed")
		t.Run("brew install tree")
	}
	fmt.Println("tree is installed")
	t.Run("tree .")
}
```

Example output when `tree` is not installed:
```
❯ go run main.go
tree is not installed
🚗 brew install tree
==> Downloading https://homebrew.bintray.com/bottles/tree-1.7.0.high_sierra.bottle.1.tar.gz
==> Pouring tree-1.7.0.high_sierra.bottle.1.tar.gz
🍺  /usr/local/Cellar/tree/1.7.0: 8 files, 114.3KB
tree is installed
🚗 tree .
.
└── main.go

0 directories, 1 file
```

### SSH Example
```go
package main

import (
	"github.com/mgenware/j7"
	"github.com/mgenware/j7/loggers"
)

func main() {
	config := &j7.SSHConfig{
		Host: "1.2.3.4",
		User: "root",
		Auth: j7.NewKeyBasedAuth("~/key.pem"),
	}

	t := j7.NewTunnel(j7.NewSSHNode(config), loggers.NewConsoleLogger())
	t.Run("pwd")
	t.Run("ls")
}
```

Sample output:
```
🚗 pwd
/root

🚗 ls
bin
build
data
```

## Windows Support
Windows is not supported because too many Unix commands are used.
