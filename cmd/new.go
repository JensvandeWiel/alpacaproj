package cmd

import (
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/JensvandeWiel/alpacaproj/templates"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var (
	database    string
	frontend    string
	packageName string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project name]",
	Short: "Creates a new project",
	Args:  cobra.ExactArgs(1),
	Run:   runNewCmd,
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&database, "database", "d", "mysql", "Database to use")
	newCmd.Flags().StringVarP(&frontend, "frontend", "f", "none", "frontend type to use, type 'none' to skip frontend, possible values: none, inertia+react, inertia+vue, inertia+svelte")
	newCmd.Flags().StringVarP(&packageName, "package", "p", "", "package name")
}

func runNewCmd(cmd *cobra.Command, args []string) {
	lvl := slog.LevelInfo
	if verbose {
		lvl = slog.LevelDebug
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: lvl,
	}))

	l.Info("Creating new project")

	proj, err := project.NewProject(args[0], database, frontend, packageName, l)
	if err != nil {
		panic(err)
	}

	//TODO check if project already exists

	err = generateProject(proj)
	if err != nil {
		panic(err)
	}
}

func generateProject(prj *project.Project) error {
	err := os.MkdirAll(prj.Path, os.ModePerm)
	if err != nil {
		return err
	}

	err = templates.BuildApp(prj)
	if err != nil {
		return err
	}

	err = templates.BuildCMD(prj)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"stores"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"models"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"facades"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"services"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"requests"}, 0755)
	if err != nil {
		return err
	}

	err = templates.BuildFlash(prj)

	if err != nil {
		return err
	}

	err = templates.BuildFrontend(prj)
	if err != nil {
		return err
	}

	err = templates.BuildHandlers(prj)
	if err != nil {
		return err
	}

	err = templates.BuildHelpers(prj)
	if err != nil {
		return err
	}

	err = templates.BuildMiddleware(prj)
	if err != nil {
		return err
	}

	err = templates.BuildMigrations(prj)
	if err != nil {
		return err
	}

	err = templates.BuildRootFiles(prj)
	if err != nil {
		return err
	}

	err = prj.SaveConfig()
	if err != nil {
		return err
	}

	prj.Logger.Info("Project created")
	return nil
}
