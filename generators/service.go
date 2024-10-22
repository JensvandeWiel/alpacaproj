package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/iancoleman/strcase"
	"strings"
)

type ServiceGenerator struct {
	camelName string
	snakeName string
	prj       *project.Project
}

func NewServiceGenerator(name string, prj *project.Project) *ServiceGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Service") {
		camelName = camelName + "Service"
	}

	return &ServiceGenerator{
		camelName: camelName,
		snakeName: strcase.ToSnake(camelName),
		prj:       prj,
	}
}

//go:embed templates/service_template.go.tmpl
var serviceTemplate string

var ErrServiceExists = errors.New("service already exists")

func (g *ServiceGenerator) Generate() error {
	g.prj.Logger.Info("Generating service: " + g.camelName)
	fileName := "services/" + g.snakeName + ".go"

	data := map[string]interface{}{
		"camelName": strings.TrimRight(g.camelName, "Service"),
	}

	err := helpers.WriteTemplateToFile(g.prj, fileName, serviceTemplate, data)
	if err != nil {
		return err
	}

	g.prj.Logger.Info("Service generated: " + g.camelName)
	return nil
}
