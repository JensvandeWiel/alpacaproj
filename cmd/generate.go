package cmd

import "github.com/spf13/cobra"

// newCmd represents the new command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates new stuff",
}

var (
	generateWorkingDir string
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().StringVarP(&generateWorkingDir, "dir", "d", ".", "working directory")
}
