package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/Omar-Arabi1/saveClip/internal/models"
	"github.com/Omar-Arabi1/saveClip/internal/utils"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the last thing in your clipboard history with a label and a priority",
	Long: `Save the last thing in your clipboard history with a label and a priority,
the priority is optional it takes numbers in this range 1 min and 3 max, but the label is
mandatory, the app will automatically save all the enteries with a creation date Y-M-D

keep in mind that all of the operations in this program will use the label for searching,
accessing, removing so I suggest choosing short one word discreptive names, For example:

saveClip save <label> --priority <priority>

keep in mind that the label and the text in your clipboard to save must be unique`,
	Run: func(cmd *cobra.Command, args []string) {
		expectedArgsNum := 1
		err := utils.GotExpectedArgs(args, expectedArgsNum)

		if err != nil {
			log.Fatal(err)
		}

		var clips []models.Clip
		clips, err = utils.ReadJson()

		if err != nil {
			log.Fatal(err)
		}

		err = clipboard.Init()

		if err != nil {
			log.Fatal(err)
		}

		var lastClipboardEnteryBytes []byte = clipboard.Read(clipboard.FmtText)
		var lastClipboardEntery string = string(lastClipboardEnteryBytes)

		label := args[0]

		err = utils.IsEmpty(label)

		if err != nil {
			log.Fatal(err)
		}

		var priority int
		priority, err = cmd.Flags().GetInt("priority")

		if err != nil {
			log.Fatal(err)
		}

		const maxPriority int = 3
		const minPriority int = 1
		if priority < minPriority {
			err = errors.New("priority is smaller than the minimum priority")
			log.Fatal(err)
		} else if priority > maxPriority {
			err = errors.New("priority is greater than the maximum priority")
			log.Fatal(err)
		}

		var currentDate string = utils.GetCurrentDate()

		var newClip models.Clip = models.Clip{
			Body:         lastClipboardEntery,
			Label:        label,
			Priority:     priority,
			CreationDate: currentDate,
		}

		for _, clip := range clips {
			if clip.Body == newClip.Body {
				err = errors.New("the text you are trying to save already exists")
				log.Fatal(err)
			} else if clip.Label == newClip.Label {
				err = errors.New("the label you are trying to save with already exists")
				log.Fatal(err)
			}
		}

		clips = append(clips, newClip)

		err = utils.WriteJson(clips)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Added The Clip successfully")
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().IntP("priority", "p", 1, "set the priority of the entery")
}
