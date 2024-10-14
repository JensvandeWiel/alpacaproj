package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
)

//go:embed sources/test_helpers/setup_db_mysql.go.tmpl
var setup_test_helpers_mysql_template []byte

//go:embed sources/test_helpers/setup_db_postgres.go.tmpl
var setup_test_helpers_postgres_template []byte

func BuildTestHelpers(prj *project.Project) error {
	prj.Logger.Debug("Generating test helpers")

	err := helpers.CreateDirectories(prj.Path, []string{"test_helpers"}, 0755)
	if err != nil {
		return err
	}

	switch prj.Database {
	case "mysql":
		err = buildTestHelpers_MySQL(prj)
		if err != nil {
			return err
		}
	case "postgres":
		err = buildTestHelpers_Postgres(prj)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildTestHelpers_MySQL(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/setup_db.go")

	tmpl, err := template.New("setup_test_helpers_mysql").Parse(string(setup_test_helpers_mysql_template))

	file, err := os.OpenFile(path.Join(prj.Path, "test_helpers", "setup_db.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	prj.Logger.Debug("Generated test_helpers/setup_db.go")
	return nil
}

func buildTestHelpers_Postgres(prj *project.Project) error {
	prj.Logger.Debug("Generating test_helpers/setup_db.go")

	tmpl, err := template.New("setup_test_helpers_postgres").Parse(string(setup_test_helpers_postgres_template))

	file, err := os.OpenFile(path.Join(prj.Path, "test_helpers", "setup_db.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	prj.Logger.Debug("Generated test_helpers/setup_db.go")
	return nil
}
