package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/iancoleman/strcase"
	"strings"
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

var ErrStoreExists = errors.New("store already exists")

func (g *StoreGenerator) Generate() error {
	g.prj.Logger.Info("Generating store: " + g.camelName)
	fileName := "stores/" + g.snakeName + ".go"

	data := map[string]interface{}{
		"camelName": g.camelName,
	}

	err := helpers.WriteTemplateToFile(g.prj, fileName, storeTemplate, data)
	if err != nil {
		return err
	}

	g.prj.Logger.Info("Store generated: " + g.camelName)
	return nil
}
