package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/test_helpers/setup_db_mysql.go.tmpl
var setupTestHelpersMySQLTemplate string

//go:embed sources/test_helpers/setup_db_postgres.go.tmpl
var setupTestHelpersPostgresTemplate string

func BuildTestHelpers(prj *project.Project) error {
	prj.Logger.Debug("Generating test helpers")

	err := helpers.CreateDirectories(prj.Path, []string{"test_helpers"}, 0755)
	if err != nil {
		return err
	}

	switch prj.Database {
	case "mysql":
		err = buildTestHelpersMySQL(prj)
		if err != nil {
			return err
		}
	case "postgres":
		err = buildTestHelpersPostgres(prj)
		if err != nil {
			return err
		}
	}

	err = buildEchoContext(prj)
	if err != nil {
		return err
	}

	return nil
}

func buildTestHelpersMySQL(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/setup_db.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "test_helpers/setup_db.go", setupTestHelpersMySQLTemplate, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated test_helpers/setup_db.go")
	return nil
}

func buildTestHelpersPostgres(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/setup_db.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "test_helpers/setup_db.go", setupTestHelpersPostgresTemplate, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated test_helpers/setup_db.go")
	return nil
}

//go:embed sources/test_helpers/echo_context_template.go.tmpl
var echoContextTemplate string

func buildEchoContext(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/echo_context.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "test_helpers/echo_context.go", echoContextTemplate, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated test_helpers/echo_context.go")
	return nil
}
