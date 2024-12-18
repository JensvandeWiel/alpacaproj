package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/iancoleman/strcase"
	"strings"
)

type FacadeGenerator struct {
	camelName string
	snakeName string
	prj       *project.Project
}

func NewFacadeGenerator(name string, prj *project.Project) *FacadeGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Facade") {
		camelName = camelName + "Facade"
	}

	return &FacadeGenerator{
		camelName: camelName,
		snakeName: strcase.ToSnake(camelName),
		prj:       prj,
	}
}

//go:embed templates/facade_template.go.tmpl
var facadeTemplate string

var ErrFacadeExists = errors.New("facade already exists")

func (g *FacadeGenerator) Generate() error {
	g.prj.Logger.Info("Generating facade: " + g.camelName)
	fileName := "facades/" + g.snakeName + ".go"

	data := map[string]interface{}{
		"camelName": g.camelName,
	}

	err := helpers.WriteTemplateToFile(g.prj, fileName, facadeTemplate, data)
	if err != nil {
		return err
	}

	g.prj.Logger.Info("Generated facade: " + g.camelName)
	return nil
}
