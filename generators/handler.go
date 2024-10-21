package generators

import (
	_ "embed"
	"errors"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"github.com/iancoleman/strcase"
	"os"
	"path"
	"strings"
)

type HandlerGenerator struct {
	camelName string
	snakeName string
	prj       *project.Project
}

func NewHandlerGenerator(name string, prj *project.Project) *HandlerGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Handler") {
		camelName = camelName + "Handler"
	}

	return &HandlerGenerator{
		camelName: camelName,
		snakeName: strcase.ToSnake(camelName),
		prj:       prj,
	}
}

//go:embed templates/handler_template.go.tmpl
var handlerTemplate string

var ErrHandlerExists = errors.New("handler already exists")

func (g *HandlerGenerator) Generate() error {
	g.prj.Logger.Info("Generating handler: " + g.camelName)
	fileName := path.Join(g.prj.Path, "handlers", g.snakeName+".go")

	err := helpers.CreateDirectories(g.prj.Path, []string{"handlers"}, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrHandlerExists
	}

	data := map[string]interface{}{
		"handlerName": g.camelName,
		"snakeName":   g.snakeName,
		"packageName": g.prj.PackageName,
	}

	err = helpers.WriteTemplateToFile(g.prj, fileName, handlerTemplate, data)
	if err != nil {
		return err
	}

	g.prj.Logger.Info("Handler generated: " + g.camelName)
	return nil
}
