package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/unisx/dr/module/cmd"
)

var (
	Version = "0.0.0"

	// GitHash value will be set during build
	GitHash = "Not provided"

	// BuildTime value will be set during build
	BuildTime = "Not provided"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "version of application",
		Long:  "version infomation for application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s \nBuildTime:%s\nGitHash:%s\n",
				Version, BuildTime, GitHash)
		},
	}

	// Register versionCmd as sub-command
	cmd.Register(versionCmd)
}
