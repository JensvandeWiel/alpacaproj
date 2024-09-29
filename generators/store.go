package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
	"text/template"
)

type StoreGenerator struct {
	camelName string
	snakeName string
	prj       *project.Project
}

func NewStoreGenerator(name string, prj *project.Project) *StoreGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Store") {
		camelName = camelName + "Store"
	}

	return &StoreGenerator{
		camelName: camelName,
		snakeName: strcase.ToSnake(camelName),
		prj:       prj,
	}
}

//go:embed templates/store_template.go.tmpl
var storeTemplate string

var ErrStoreExists = errors.New("Store already exists")

func (g *StoreGenerator) Generate() error {
	g.prj.Logger.Info("Generating Store: " + g.camelName)
	fileName := g.prj.Path + "/stores/" + g.snakeName + ".go"

	tmpl, err := template.New("Store").Parse(storeTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"stores"}, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrStoreExists
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"camelName": g.camelName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	g.prj.Logger.Info("Generated Store: " + g.camelName)
	return nil
}
