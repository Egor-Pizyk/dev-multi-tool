package core

import (
	"dev-multi-tool/tools"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var UUIDCommand = &cobra.Command{
	Use:   "uuid",
	Short: "Generate uuid",
	Run: func(cmd *cobra.Command, args []string) {
		var uuid = tools.GetUUID()
		err := clipboard.WriteAll(uuid.String())
		if err == nil {
			fmt.Println("Already in clipboard.")
		}
		fmt.Println(uuid)
	},
}
