package cmd

import (
	"github.com/JensvandeWiel/alpacaproj/generators"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/spf13/cobra"
)

var generateServiceCmd = &cobra.Command{
	Use:     "service [Service name]",
	Short:   "Generates a new Service",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"Service"},
	RunE:    runGenerateServiceCmd,
}

func init() {
	generateCmd.AddCommand(generateServiceCmd)
}

func runGenerateServiceCmd(cmd *cobra.Command, args []string) error {
	prj, err := project.LoadProject(generateWorkingDir, verbose)
	if err != nil {
		return err
	}

	g := generators.NewServiceGenerator(args[0], prj)
	err = g.Generate()
	if err != nil {
		return err
	}

	return nil
}
