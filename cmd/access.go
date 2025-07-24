package cmd

import (
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
		var expectedArgsNum int = 0
		err := utils.GotExpectedArgs(args, expectedArgsNum)

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

		for _, clip := range slices.Backward(clips) {
			clipboard.Write(clipboard.FmtText, []byte(clip.Body))
			fmt.Printf("copied %s into your clipboard\n", clip.Body)
			time.Sleep(30 * time.Millisecond)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(accessCmd)
}
