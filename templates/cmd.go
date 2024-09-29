package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
)

//go:embed sources/cmd/server/database_mysql.go.tmpl
var server_database_template []byte

// buildCMDServer_DatabaseMySQL generates the cmd/server/database_mysql.go file
func buildCMDServer_DatabaseMySQL(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/database_mysql.go")
	if prj.Database != project.MySQL {
		prj.Logger.Debug("Database is not MySQL, skipping")
		return nil
	}

	tmpl, err := template.New("server_database_mysql").Parse(string(server_database_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"cmd/server"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "cmd/server/database_mysql.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, nil)

}

//go:embed sources/cmd/server/database_postgres.go.tmpl
var server_database_postgres_template []byte

func buildCMDServer_DatabasePostgres(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/database_postgres.go")
	if prj.Database != project.Postgres {
		prj.Logger.Debug("Database is not Postgres, skipping")
		return nil
	}

	tmpl, err := template.New("server_database_postgres").Parse(string(server_database_postgres_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"cmd/server"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "cmd/server/database_postgres.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, nil)
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
var server_goose_template []byte

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

	tmpl, err := template.New("server_goose").Parse(string(server_goose_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"cmd/server"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "cmd/server/goose.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]interface{}{
		"dialect":     parseGooseDialect(prj.Database),
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
}

//go:embed sources/cmd/server/main.go.tmpl
var server_main_template []byte

// buildCMDServer_Main generates the cmd/server/main.go file
func buildCMDServer_Main(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/main.go")

	tmpl, err := template.New("server_main").Parse(string(server_main_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"cmd/server"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "cmd/server/main.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
		"projectName": prj.Name,
		"hasFrontend": prj.HasFrontend,
		"database":    prj.Database,
	}

	return tmpl.Execute(file, data)
}

//go:embed sources/cmd/server/serve.go.tmpl
var server_serve_template []byte

// buildCMDServer_Serve generates the cmd/server/serve.go file
func buildCMDServer_Serve(prj *project.Project) error {
	prj.Logger.Debug("Generating cmd/server/serve.go")

	tmpl, err := template.New("server_serve").Parse(string(server_serve_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"cmd/server"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "cmd/server/serve.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
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
