# lun [Working in progress]
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
	"github.com/mgenware/lun/runners"
)

func main() {
	n := lun.NewLocalNode()
	r := runners.NewConsoleRunner()

	_, err := exec.LookPath("tree")
	if err != nil {
		fmt.Println("tree is not installed")
		r.Run(n, "brew install tree")
	}
	fmt.Println("tree is installed")
	r.Run(n, "tree .")
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
	"log"

	"github.com/mgenware/lun"
)

func main() {
	config := &lun.SSHConfig{
		Host: "123.4.5.6",
		User: "root",
		Auth: lun.MustNewKeyBasedAuth("./key.pem"),
	}

	node := lun.MustNewSSHNode(config)
	output, err := node.Exec("pwd")
	if err != nil {
		panic(err)
	}
	log.Print(string(output))
}
```

Sample output:
```
2018/10/02 01:04:31 /root
```
