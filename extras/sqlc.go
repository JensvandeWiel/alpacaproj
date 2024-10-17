package extras

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
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
var sqlc_template []byte

func buildSQLCConfig(prj *project.Project) error {
	prj.Logger.Debug("Generating sqlc.yaml")

	templ, err := template.New("sqlc").Parse(string(sqlc_template))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "sqlc.yaml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	dbDriver := prj.Database
	if dbDriver == project.Postgres {
		dbDriver = "postgresql"
	}

	data := map[string]interface{}{
		"databaseDriver": dbDriver,
	}

	err = templ.Execute(file, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Generated sqlc.yaml")

	return nil
}
