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
	Increments the pth portion of a SmeVer in the specified package.json
	$ terrafold bumpsemver ./package.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: terrafold bumpsemver <path to package.json>")
			return
		}
		err := deploy.BumpPackage(args[0])
		if err != nil {
			fmt.Println("ERROR Bumping SemVer", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bumpsemverCmd)
}
