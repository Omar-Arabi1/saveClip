package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "saveClip",
	Short: "a clipboard manager to save only imporant things to you",
	Version: "0.1.0",
	Long: `a clipboard manager tool to save only important things to you
by running the save command you will save the latest thing in your clipboard
you could use the --help option for more info when you need it, example:

saveClip <command> <option>`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
