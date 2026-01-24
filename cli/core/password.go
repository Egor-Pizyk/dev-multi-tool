package core

import (
	"dev-multi-tool/common"
	"dev-multi-tool/tools"
	"fmt"

	"github.com/atotto/clipboard"
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
			fmt.Println(common.ErrorsEnum[common.SomethingWrong])
		}
		if err := clipboard.WriteAll(password); err != nil {
			fmt.Println(common.ErrorsEnum[common.SomethingWrong])
			return
		}
		fmt.Println(password)
	},
}

func init() {
	PasswordCommand.Flags().IntVarP(&size, "size", "s", 32, "Password length")
	PasswordCommand.Flags().BoolVar(&isNumbersInUsage, "num", false, "Are numbers should used")
}
