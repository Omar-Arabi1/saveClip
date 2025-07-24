package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "See all your enteries inside the list",
	Long: `see all your enteries inside the list, by default the app
will print them out by the order they are saved in the list you could 
override this by giving the command --priority option which takes in highest/lowest 
and sorts by priority or the --date option which takes in oldest/newest and sort by date
For example:

saveClip view --priority highest`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("view called")
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
