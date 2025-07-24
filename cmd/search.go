package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Fuzzy searches through your list by the given query",
	Long: `fuzzy searches through your entire list with a query to search with
it will print out the labels with that query inside

the query itself will be part of the label name not the copied text itself, For example:

saveClip search <query>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
