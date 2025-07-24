package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an item from your list",
	Long: `removes an item from your list it takes in a label to an item in the list to remove
it has two options --all removes all entries at once --remove-priority removes all enteries with
the same priority. For example:

saveClip remove <label>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
