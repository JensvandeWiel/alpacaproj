package cmd

import (
	"github.com/JensvandeWiel/alpacaproj/generators"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/spf13/cobra"
)

var generateRequestCmd = &cobra.Command{
	Use:     "request [Request name]",
	Short:   "Generates a new Request",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"Request"},
	RunE:    runGenerateRequestCmd,
}

func init() {
	generateCmd.AddCommand(generateRequestCmd)
}

func runGenerateRequestCmd(cmd *cobra.Command, args []string) error {
	prj, err := project.LoadProject(generateWorkingDir, verbose)
	if err != nil {
		return err
	}

	g := generators.NewRequestGenerator(args[0], prj)
	err = g.Generate()
	if err != nil {
		return err
	}

	return nil
}
