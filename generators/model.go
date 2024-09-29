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

type ModelGenerator struct {
	camelName string
	snakeName string
	prj       *project.Project
}

func NewModelGenerator(name string, prj *project.Project) *ModelGenerator {
	camelName := strcase.ToCamel(name)

	if !strings.HasSuffix(camelName, "Model") {
		camelName = camelName + "Model"
	}

	return &ModelGenerator{
		camelName: camelName,
		snakeName: strcase.ToSnake(camelName),
		prj:       prj,
	}
}

//go:embed templates/model_template.go.tmpl
var modelTemplate string

var ErrModelExists = errors.New("model already exists")

func (g *ModelGenerator) Generate() error {
	g.prj.Logger.Info("Generating model: " + g.camelName)
	fileName := g.prj.Path + "/models/" + g.snakeName + ".go"

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
		"camelName": strings.TrimRight(g.camelName, "Model"),
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	g.prj.Logger.Info("Generated model: " + g.camelName)

	return nil
}
