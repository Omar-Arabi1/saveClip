package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/Omar-Arabi1/saveClip/internal/models"
	"github.com/Omar-Arabi1/saveClip/internal/utils"
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
		var expectedArgsNum int = 1
		err := utils.GotExpectedArgs(args, expectedArgsNum)

		if err != nil {
			log.Fatal(err)
		}

		label := args[0]

		var clips []models.Clip
		clips, err = utils.ReadJson()

		if err != nil {
			log.Fatal(err)
		}

		err = utils.AreClipsEmpty(clips)

		if err != nil {
			log.Fatal(err)
		}

		var newClips []models.Clip

		for _, clip := range clips {
			if clip.Label != label {
				newClips = append(newClips, clip)
			}
		}

		if len(newClips) == len(clips) {
			err = errors.New("didn't find any clip with the given label")
			log.Fatal(err)
		}

		err = utils.WriteJson(newClips)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Removed the entery successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
