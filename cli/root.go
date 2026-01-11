package cli

import (
	"dev-multi-tool/cli/core"
	"dev-multi-tool/cli/sub"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "dmt",
	Short:        "Dev Multi Tool",
	SilenceUsage: true,
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(core.UUIDCommand)
	rootCmd.AddCommand(core.PasswordCommand)
	rootCmd.AddCommand(core.PasswordCommand)
	rootCmd.AddCommand(sub.VersionCommand)
}
