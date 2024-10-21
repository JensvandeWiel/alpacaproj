package helpers

import "github.com/JensvandeWiel/alpacaproj/project"

func RunGoTidy(prj *project.Project) error {
	prj.Logger.Debug("Running go mod tidy")

	err := RunCommand(prj.Path, true, "go", "mod", "tidy")
	if err != nil {
		return err
	}

	prj.Logger.Debug("Ran go mod tidy")

	return nil
}
