package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/iancoleman/strcase"
	"strings"
)

type RequestGenerator struct {
	camelName string
	snakeName string
	prj       *project.Project
}

func NewRequestGenerator(name string, prj *project.Project) *RequestGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Request") {
		camelName = camelName + "Request"
	}

	return &RequestGenerator{
		camelName: camelName,
		snakeName: strcase.ToSnake(camelName),
		prj:       prj,
	}
}

//go:embed templates/request_template.go.tmpl
var requestTemplate string

var ErrRequestExists = errors.New("request already exists")

func (g *RequestGenerator) Generate() error {
	g.prj.Logger.Info("Generating request: " + g.camelName)
	fileName := "requests/" + g.snakeName + ".go"

	data := map[string]interface{}{
		"camelName": g.camelName,
	}

	err := helpers.WriteTemplateToFile(g.prj, fileName, requestTemplate, data)
	if err != nil {
		return err
	}

	g.prj.Logger.Info("Request generated: " + g.camelName)
	return nil
}
