package cmd

import (
	"github.com/JensvandeWiel/alpacaproj/generators"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/spf13/cobra"
)

var generateStoreCmd = &cobra.Command{
	Use:     "store [Store name]",
	Short:   "Generates a new Store",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"Store"},
	RunE:    runGenerateStoreCmd,
}

func init() {
	generateCmd.AddCommand(generateStoreCmd)
}

func runGenerateStoreCmd(cmd *cobra.Command, args []string) error {
	prj, err := project.LoadProject(generateWorkingDir, verbose)
	if err != nil {
		return err
	}

	g := generators.NewStoreGenerator(args[0], prj)
	err = g.Generate()
	if err != nil {
		return err
	}

	return nil
}
