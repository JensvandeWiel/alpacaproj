package templates

import (
	"embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"path"
)

func BuildMigrations(prj *project.Project) error {
	prj.Logger.Debug("Generating migrations")

	err := helpers.CreateDirectories(prj.Path, []string{"migrations"}, 0755)
	if err != nil {
		return err
	}

	switch prj.Database {
	case "mysql":
		err = buildMigrations_MySQL(prj)
		if err != nil {
			return err
		}
	case "postgres":
		err = buildMigrations_Postgres(prj)
		if err != nil {
			return err
		}
	}

	return nil
}

//go:embed sources/mysql/migrations
var mysql_migrations_template embed.FS

func buildMigrations_MySQL(prj *project.Project) error {
	prj.Logger.Debug("Generating migrations for MySQL")

	err := helpers.CreateDirectories(prj.Path, []string{"migrations"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(mysql_migrations_template, "sources/mysql/migrations", path.Join(prj.Path, "migrations"))
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/postgres/migrations
var postgres_migrations_template embed.FS

func buildMigrations_Postgres(prj *project.Project) error {
	prj.Logger.Debug("Generating migrations for Postgres")

	err := helpers.CreateDirectories(prj.Path, []string{"migrations"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(postgres_migrations_template, "sources/postgres/migrations", path.Join(prj.Path, "migrations"))
	if err != nil {
		return err
	}

	return nil
}
