package cmd

import (
	"fmt"

	"github.com/bartalcorn/terrafold/pkg/scaffold"
	"github.com/spf13/cobra"
)

// newProlifeCmd represents the list command
var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "scaffold a new lambda from a profile.json",
	Long: `
	Use the named profile to scaffold a new lambda!
	$ terrafold scaffold name-of-new-profile`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: terrafold scaffold <profile-name>")
			return
		}
		scaffold.Do(args[0])
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}
