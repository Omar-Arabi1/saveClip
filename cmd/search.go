package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/Omar-Arabi1/saveClip/internal/models"
	"github.com/Omar-Arabi1/saveClip/internal/utils"
	"github.com/lithammer/fuzzysearch/fuzzy"
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
		filterDate, err := cmd.Flags().GetString("filter")
		
		if err != nil {
			log.Fatal(err)
		}
		
		var expectedArgsNum int = 1
		err = utils.GotExpectedArgs(args, expectedArgsNum)

		if err != nil {
			log.Fatal(err)
		}

		query := args[0]

		var clips []models.Clip
		clips, err = utils.ReadJson()

		if err != nil {
			log.Fatal(err)
		}

		err = utils.AreClipsEmpty(clips)

		if err != nil {
			log.Fatal(err)
		}

		var clipsText []string

		for _, clip := range clips {
			if filterDate == "" {
				clipsText = append(clipsText, clip.Label)
			} else if clip.CreationDate == filterDate {
				clipsText = append(clipsText, clip.Label)
			}
		}

		matches := fuzzy.Find(query, clipsText)

		if len(matches) == 0 {
			err = errors.New("no clipboard label matched your query")
			log.Fatal(err)
		}
		
		for _, clip := range clips {
			for _, match := range matches {
				if clip.Label == match {
					fmt.Printf("Label: %s - Has Entery: %s\n", clip.Label, clip.Body)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("filter", "f", "", "filter the results only if they were created at the given date YYYY-MM-DD")
}
