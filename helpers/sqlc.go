package helpers

import "github.com/JensvandeWiel/alpacaproj/project"

func GenerateSQLCDefinitions(prj *project.Project) error {
	prj.Logger.Debug("Generating SQLC definitions")

	err := RunCommand(prj.Path, true, "sqlc", "generate")
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated SQLC definitions")

	return nil
}
