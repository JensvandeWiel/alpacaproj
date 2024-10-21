package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/cmd/server/database_mysql.go.tmpl
var serverDatabaseMySQLTemplate string

// buildCMDServer_DatabaseMySQL generates the cmd/server/database_mysql.go file
func buildCMDServer_DatabaseMySQL(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/database_mysql.go")
	if prj.Database != project.MySQL {
		prj.Logger.Debug("Database is not MySQL, skipping")
		return nil
	}

	err := helpers.WriteTemplateToFile(prj, "cmd/server/database_mysql.go", serverDatabaseMySQLTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/cmd/server/database_postgres.go.tmpl
var serverDatabasePostgresTemplate string

// buildCMDServer_DatabasePostgres generates the cmd/server/database_postgres.go file
func buildCMDServer_DatabasePostgres(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/database_postgres.go")
	if prj.Database != project.Postgres {
		prj.Logger.Debug("Database is not Postgres, skipping")
		return nil
	}

	err := helpers.WriteTemplateToFile(prj, "cmd/server/database_postgres.go", serverDatabasePostgresTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

func buildCMDServer_Database(prj *project.Project) error {
	switch prj.Database {
	case project.MySQL:
		return buildCMDServer_DatabaseMySQL(prj)
	case project.Postgres:
		return buildCMDServer_DatabasePostgres(prj)
	default:
		panic("unknown database driver")
	}
}

//go:embed sources/cmd/server/goose.go.tmpl
var serverGooseTemplate string

func parseGooseDialect(database project.DatabaseDriver) string {
	switch database {
	case "mysql":
		return "mysql"
	case "postgres":
		return "postgres"
	default:
		return ""
	}
}

// buildCMDServer_Goose generates the cmd/server/goose.go file
func buildCMDServer_Goose(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/goose.go")

	data := map[string]interface{}{
		"dialect":     parseGooseDialect(prj.Database),
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "cmd/server/goose.go", serverGooseTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/cmd/server/main.go.tmpl
var serverMainTemplate string

// buildCMDServer_Main generates the cmd/server/main.go file
func buildCMDServer_Main(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/main.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"projectName": prj.Name,
		"hasFrontend": prj.HasFrontend,
		"database":    prj.Database,
	}

	err := helpers.WriteTemplateToFile(prj, "cmd/server/main.go", serverMainTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/cmd/server/serve.go.tmpl
var serverServeTemplate string

// buildCMDServer_Serve generates the cmd/server/serve.go file
func buildCMDServer_Serve(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/serve.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "cmd/server/serve.go", serverServeTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

func BuildCMD(prj *project.Project) error {
	err := buildCMDServer_Database(prj)
	if err != nil {
		return err
	}

	err = buildCMDServer_Goose(prj)
	if err != nil {
		return err
	}

	err = buildCMDServer_Main(prj)
	if err != nil {
		return err
	}

	err = buildCMDServer_Serve(prj)
	if err != nil {
		return err
	}

	return nil
}
