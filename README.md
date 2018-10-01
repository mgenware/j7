# lun [Working in progress]
Shell scripts in Go

## Installation
```sh
go get github.com/mgenware/lun
```

## Examples
### SSH
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
