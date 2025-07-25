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
		removeAll, err := cmd.Flags().GetBool("all")

		if err != nil {
			log.Fatal(err)
		}

		if removeAll {
			err = utils.WriteJson([]models.Clip{})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Removed all clips successfully")
			return
		}

		var removePriority int
		removePriority, err = cmd.Flags().GetInt("remove-priority")

		if err != nil {
			log.Fatal(err)
		}

		var label string

		if removePriority == 0 {
			var expectedArgsNum int = 1
			err = utils.GotExpectedArgs(args, expectedArgsNum)

			if err != nil {
				log.Fatal(err)
			}

			label = args[0]
		} else {
			var expectedArgsNum int = 0
			err = utils.GotExpectedArgs(args, expectedArgsNum)

			if err != nil {
				log.Fatal(err)
			}
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

		var newClips []models.Clip

		for _, clip := range clips {
			if clip.Label != label && removePriority == 0{
				newClips = append(newClips, clip)
			} else if clip.Priority != removePriority && removePriority != 0 {
				newClips = append(newClips, clip)
			}
		}

		if len(newClips) == len(clips) {
			err = errors.New("didn't remove any clips")
			log.Fatal(err)
		}

		err = utils.WriteJson(newClips)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("The removal was successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolP("all", "a", false, "remove all clips at once")
	removeCmd.Flags().IntP("remove-priority", "p", 0, "remove all clips that have the given priority")
}
