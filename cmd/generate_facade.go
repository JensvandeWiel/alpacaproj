package cmd

import (
	"github.com/JensvandeWiel/alpacaproj/generators"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/spf13/cobra"
)

var generateFacadeCmd = &cobra.Command{
	Use:     "facade [Facade name]",
	Short:   "Generates a new Facade",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"Facade"},
	RunE:    runGenerateFacadeCmd,
}

func init() {
	generateCmd.AddCommand(generateFacadeCmd)
}

func runGenerateFacadeCmd(cmd *cobra.Command, args []string) error {
	prj, err := project.LoadProject(generateWorkingDir, verbose)
	if err != nil {
		return err
	}

	g := generators.NewFacadeGenerator(args[0], prj)
	err = g.Generate()
	if err != nil {
		return err
	}

	return nil
}
