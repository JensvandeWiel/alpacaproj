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
	fileName := g.prj.Path + "/facades/" + g.snakeName + ".go"

	tmpl, err := template.New("facade").Parse(facadeTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"facades"}, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrFacadeExists
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
	g.prj.Logger.Info("Generated facade: " + g.camelName)
	return nil
}
