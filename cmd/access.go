package cmd

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/Omar-Arabi1/saveClip/internal/models"
	"github.com/Omar-Arabi1/saveClip/internal/utils"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
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
		label, err := cmd.Flags().GetString("label")

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

		err = clipboard.Init()
		if err != nil {
			log.Fatal(err)
		}

		var copiedClipBody string

		for _, clip := range slices.Backward(clips) {
			if label == "" {
				copiedClipBody = clip.Body
				break
			} else if clip.Label == label {
				copiedClipBody = clip.Body
				break
			}
		}

		if copiedClipBody == "" {
			err = errors.New("could not find a clip with the given label")
			log.Fatal(err)
		}

		clipboard.Write(clipboard.FmtText, []byte(copiedClipBody))
		fmt.Printf("copied %s into your clipboard\n", copiedClipBody)
		time.Sleep(30 * time.Millisecond)
	},
}

func init() {
	rootCmd.AddCommand(accessCmd)
	accessCmd.Flags().StringP("label", "l", "", "choose a different clip to access by providing its label")
}
