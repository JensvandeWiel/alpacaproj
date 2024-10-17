package cmd

import (
	"github.com/JensvandeWiel/alpacaproj/generators"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/spf13/cobra"
)

var generateModelCmd = &cobra.Command{
	Use:     "model [Model name]",
	Short:   "Generates a new Model",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"Model"},
	RunE:    runGenerateModelCmd,
}

var (
	addExtras bool
)

func init() {
	generateCmd.AddCommand(generateModelCmd)
	generateModelCmd.Flags().BoolVarP(&addExtras, "extras", "e", false, "Add extras to the model: store and tests")
}

func runGenerateModelCmd(cmd *cobra.Command, args []string) error {
	prj, err := project.LoadProject(generateWorkingDir, verbose)
	if err != nil {
		return err
	}

	g := generators.NewModelGenerator(args[0], prj, addExtras)
	err = g.Generate()
	if err != nil {
		return err
	}

	return nil
}
