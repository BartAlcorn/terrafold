package cmd

import (
	"fmt"

	"github.com/bartalcorn/terrafold/pkg/efs"
	"github.com/spf13/cobra"
)

// newProlifeCmd represents the list command
var dumptemplatesCmd = &cobra.Command{
	Use:   "dumptemplates",
	Short: "dumptemplates",
	Long: `
	Dumps the embedded templates to folder named terrafoldTemplates
	$ terrafold dumptemplates`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: terrafold dumptemplates")
			return
		}
		efs.DumpTemplates()
	},
}

func init() {
	rootCmd.AddCommand(dumptemplatesCmd)
}
