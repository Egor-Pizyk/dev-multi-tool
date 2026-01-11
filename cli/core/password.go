package core

import (
	"dev-multi-tool/tools"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	size int
	//isNumbersInUsage bool
	//isSymbolsInUsage bool
)

var PasswordCommand = &cobra.Command{
	Use:   "pass",
	Short: "Generate password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tools.GetPassword(size))
	},
}

func init() {
	PasswordCommand.Flags().IntVarP(&size, "size", "s", 32, "Password length")
	//PasswordCommand.Flags().BoolVar(&isNumbersInUsage, "num", false, "Are numbers should used")
	//PasswordCommand.Flags().BoolVar(&isSymbolsInUsage, "symbol", false, "Are symbols should used")
}
