package cmd

import (
	"fmt"

	"github.com/bartalcorn/terrafold/pkg/scaffold"
	"github.com/spf13/cobra"
)

// newProlifeCmd represents the list command
var newProfileCmd = &cobra.Command{
	Use:   "newprofile",
	Short: "newprofile",
	Long: `
	Creates a new named profile
	$ terrafold newProfile name-of-new-profile`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("usage: terrafold newprofile <new-profile-name>")
			return
		}
		scaffold.NewProfile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(newProfileCmd)
}
