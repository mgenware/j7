# lun

[![Build Status](https://travis-ci.org/mgenware/lun.svg?branch=master)](http://travis-ci.org/mgenware/lun)

Shell scripting in Go

## Installation
```sh
go get github.com/mgenware/lun
```

## Examples
Checking if a command is installed and performing an install command if necessary. (assume macOS with homebrew installed)
```go
package main

import (
	"fmt"
	"os/exec"

	"github.com/mgenware/lun"
	"github.com/mgenware/lun/loggers"
)

func main() {
	w := lun.NewNodeWrapper(lun.NewLocalNode(), loggers.NewConsoleLogger())

	_, err := exec.LookPath("tree")
	if err != nil {
		w.Logger().Log(lun.LogLevelError, "tree is not installed")
		w.Run("brew install tree")
	}
	fmt.Println("tree is installed")
	w.Run("tree .")
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
	"github.com/mgenware/lun"
	"github.com/mgenware/lun/loggers"
)

func main() {
	config := &lun.SSHConfig{
		Host: "1.2.3.4",
		User: "root",
		Auth: lun.NewKeyBasedAuth("~/key.pem"),
	}

	w := lun.NewNodeWrapper(lun.MustNewSSHNode(config), loggers.NewConsoleLogger())
	w.Run("pwd")
	w.Run("ls")
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
