package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
	"text/template"
)

type ModelGenerator struct {
	camelName         string
	addExtras         bool
	snakeName         string
	name              string
	camelNameNoSuffix string
	prj               *project.Project
}

func NewModelGenerator(name string, prj *project.Project, addExtras bool) *ModelGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Model") {
		camelName = camelName + "Model"
	}

	return &ModelGenerator{
		camelName:         camelName,
		addExtras:         addExtras,
		name:              strings.ToLower(strings.TrimSuffix(name, "Model")),
		camelNameNoSuffix: strings.TrimSuffix(camelName, "Model"),
		snakeName:         strcase.ToSnake(camelName),
		prj:               prj,
	}
}

//go:embed templates/model_template.go.tmpl
var modelTemplate string

var ErrModelExists = errors.New("model already exists")

func (g *ModelGenerator) Generate() error {
	g.prj.Logger.Info("Generating model: " + g.camelName)
	fileName := g.prj.Path + "/models/" + g.name + ".go"

	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"models"}, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrModelExists
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"camelNameNoSuffix": g.camelNameNoSuffix,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	if g.addExtras {
		g.prj.Logger.Debug("Generating extras")
		err = g.generateExtras()
		if err != nil {
			return err
		}
		g.prj.Logger.Info("Generated extras")
	}

	g.prj.Logger.Info("Generated model: " + g.camelName)

	return nil
}

//go:embed templates/extra_model/store_mysql.go.tmpl
var storeMysqlTemplate string

//go:embed templates/extra_model/store_postgres.go.tmpl
var storePostgresTemplate string

func (g *ModelGenerator) generateMysqlStore() error {
	tmpl, err := template.New("user_store_mysql").Parse(storeMysqlTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"stores"}, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := g.prj.Path + "/stores/" + g.name + "_store.go"

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrStoreExists
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	plur := pluralize.NewClient()

	data := map[string]interface{}{
		"camelNameNoSuffix": g.camelNameNoSuffix,
		"lowName":           g.name,
		"packageName":       g.prj.PackageName,
		"pluralLowName":     plur.Plural(g.name),
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func (g *ModelGenerator) generatePostgresStore() error {
	tmpl, err := template.New("user_store_postgres").Parse(storePostgresTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"stores"}, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := g.prj.Path + "/stores/" + g.name + "_store.go"

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrStoreExists
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	plur := pluralize.NewClient()

	data := map[string]interface{}{
		"camelNameNoSuffix": g.camelNameNoSuffix,
		"lowName":           g.name,
		"packageName":       g.prj.PackageName,
		"pluralLowName":     plur.Plural(g.name),
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}

//go:embed templates/extra_model/store_test.go.tmpl
var storeTestTemplate string

func (g *ModelGenerator) generateExtras() error {
	switch g.prj.Database {
	case "mysql":
		err := g.generateMysqlStore()
		if err != nil {
			return err
		}
	case "postgres":
		err := g.generatePostgresStore()
		if err != nil {
			return err
		}
	default:
		return errors.New("database not supported")
	}

	tmpl, err := template.New("store_test").Parse(storeTestTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"stores"}, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := g.prj.Path + "/stores/" + g.name + "_store_test.go"

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrStoreExists
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	plur := pluralize.NewClient()

	data := map[string]interface{}{
		"camelNameNoSuffix": g.camelNameNoSuffix,
		"lowName":           g.name,
		"packageName":       g.prj.PackageName,
		"pluralLowName":     plur.Plural(g.name),
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
