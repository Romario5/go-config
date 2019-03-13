# Configuration reading utility

To import in your project:
```
go get github.com/romario5/go-config
```

Config file example:
```ini
# Network
Port = 5000
MaxConnections = 1000
UseSSL = True

# Security
SecretToken = qwerty123456
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
	maxConnections := config.GetUint64("MaxConnections", 100)
	useSSL := config.GetBool("UseSSL", false)
	secretToken := config.GetString("SecretToken", "")
}
```

Available values for bool type that will be parsed to `true`: True, true, 1, Yes, yes.
Other values will be parsed to `false`.  