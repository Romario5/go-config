# Configuration reading utility

To import in your project:
```
go get github.com/romario5/go-config
```


Usage:
```go
package main

import (
	config "github.com/romario5/go-config"
	"fmt"
)

func main() {
	// Read configuration file.
	readProps := config.LoadFile("./general.conf")
	fmt.Println(readProps, "properties read.")
	
	// Lets get some properties:
	port := config.GetString("Port", "80")
}
```