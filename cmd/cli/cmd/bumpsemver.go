package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bartalcorn/terrafold/pkg/deploy"
)

// newProlifeCmd represents the list command
var bumpsemverCmd = &cobra.Command{
	Use:   "bumpsemver",
	Short: "bumpsemver <name of package.json>",
	Long: `
	!!  This command is being depreciated in favor of bumppatch  !!
	Increments the patch portion of a SemVer in the specified package.json
	$ terrafold bumpsemver ./package.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: terrafold bumpsemver <path to package.json>")
			return
		}
		err := deploy.BumpPatch(args[0])
		if err != nil {
			fmt.Println("ERROR Bumping SemVer", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bumpsemverCmd)
}
