package extras

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

func BuildSQLC(prj *project.Project) error {
	prj.Logger.Debug("Building extra SQLC")
	err := buildSQLCConfig(prj)

	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/sqlc/sqlc.yaml.tmpl
var sqlcTemplate string

func buildSQLCConfig(prj *project.Project) error {
	prj.Logger.Debug("Generating sqlc.yaml")

	dbDriver := prj.Database
	if dbDriver == project.Postgres {
		dbDriver = "postgresql"
	}

	data := map[string]interface{}{
		"databaseDriver": dbDriver,
	}

	err := helpers.WriteTemplateToFile(prj, "sqlc.yaml", sqlcTemplate, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated sqlc.yaml")

	return nil
}
