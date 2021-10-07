package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bartalcorn/terrafold/pkg/deploy"
)

// newProlifeCmd represents the list command
var bumpminorCmd = &cobra.Command{
	Use:   "bumpminor",
	Short: "bumpminor <name of package.json>",
	Long: `
	Increments the minor portion of a SemVer in the package.json file in the current folder.
	$ terrafold bumpminor`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deploy.BumpMinor()
		if err != nil {
			fmt.Println("ERROR Bumping SemVer", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bumpminorCmd)
}
