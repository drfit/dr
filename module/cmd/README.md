# Dr.cmd
cmd module is used in build Dr. command.

### Usage:
```go
package main

import (
    "github.com/unisx/dr/module/cmd"
)

func main() {
    // Setup root cli command of application
	cmd.Setup(
		"dr",               // command name
		"dr. command",      // command short describe
		"dr. command",      // command long describe
	)

	// Execute start application
	cmd.Execute()
}
```