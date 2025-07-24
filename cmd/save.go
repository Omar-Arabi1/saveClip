package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the last thing in your clipboard history with a label and a priority",
	Long: `Save the last thing in your clipboard history with a label and a priority, 
the priority is optional it takes numbers in this range 1 min and 3 max, but the label is
mandatory, the app will automatically save all the enteries with a creation date Y-M-D

keep in mind that all of the operations in this program will use the label for searching,
accessing, removing so I suggest choosing short one word discreptive names, For example:

saveClip save <label> --priority <priority>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("save called")
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
