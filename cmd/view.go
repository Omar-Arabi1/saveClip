package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/Omar-Arabi1/saveClip/internal/models"
	"github.com/Omar-Arabi1/saveClip/internal/utils"
	"github.com/spf13/cobra"
)

// TODO: implement the two sort options

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
		sortBy, err := cmd.Flags().GetString("priority")

		if err != nil {
			log.Fatal(err)
		}

		var expectedArgsNum int = 0
		err = utils.GotExpectedArgs(args, expectedArgsNum)

		if err != nil {
			log.Fatal(err)
		}

		var clips []models.Clip
		clips, err = utils.ReadJson()

		if err != nil {
			log.Fatal(err)
		}

		err = utils.AreClipsEmpty(clips)

		if err != nil {
			log.Fatal(err)
		}

		switch sortBy {
		case "":
			break
		case "highest":
			sortHighest := true
			utils.SortPriority(clips, sortHighest)
		case "lowest":
			sortHighest := false
			utils.SortPriority(clips, sortHighest)
		default:
			errMessage := fmt.Sprintf("expected highest/lowest got %s", sortBy)
			err = errors.New(errMessage)
			log.Fatal(err)
		}

		for _, clip := range clips {
			fmt.Printf("Created at: %s - Label: %s - content: %s - Priority: %d\n", clip.CreationDate, clip.Label, clip.Body, clip.Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().StringP("priority", "p", "", "see the list sorted by priority which takes in either a highest or lowest")
}
