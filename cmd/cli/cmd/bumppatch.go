package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bartalcorn/terrafold/pkg/deploy"
)

// newProlifeCmd represents the list command
var bumppatchCmd = &cobra.Command{
	Use:   "bumppatch",
	Short: "bumppatch <name of package.json>",
	Long: `
	Increments the patch portion of a SemVer in the specified package.json
	$ terrafold bumppatch ./package.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: terrafold bumppatch <path to package.json>")
			return
		}
		err := deploy.BumpPatch(args[0])
		if err != nil {
			fmt.Println("ERROR Bumping SemVer", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bumppatchCmd)
}
