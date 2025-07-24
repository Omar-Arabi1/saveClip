package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "access one of your enteries inside the list",
	Long: `access one of your enteries inside the list access as in put it in 
your clipboard history, by default the app will access the last item you entered
you could override this by choosing the --label option and giving it a label you want to 
access

saveClip remove --label <label>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("access called")
	},
}

func init() {
	rootCmd.AddCommand(accessCmd)
}
