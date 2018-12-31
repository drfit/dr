package main

import (
	"github.com/unisx/dr/module/cmd"
	_ "github.com/unisx/dr/module/drctl/version"
	_ "github.com/unisx/dr/module/openapi/cmd"
)

func main() {
	// Setup root cli command of application
	cmd.Setup(
		"drctl",                               // command name
		"drctl is a usefull cli tool for Dr.", // command short describe
		"drctl is a usefull cli tool for Dr.", // command long describe
	)

	// Execute start application
	cmd.Execute()
}
