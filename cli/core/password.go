package core

import (
	"dev-multi-tool/tools"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	size             int
	isNumbersInUsage bool
)

var PasswordCommand = &cobra.Command{
	Use:   "pass",
	Short: "Generate password",
	Run: func(cmd *cobra.Command, args []string) {
		var password, err = tools.GetPassword(size, isNumbersInUsage)
		if err != nil {
			fmt.Println("Something went wrong...")
		}
		fmt.Println(password)
	},
}

func init() {
	PasswordCommand.Flags().IntVarP(&size, "size", "s", 32, "Password length")
	PasswordCommand.Flags().BoolVar(&isNumbersInUsage, "num", false, "Are numbers should used")
}
