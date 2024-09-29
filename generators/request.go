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
	fileName := g.prj.Path + "/requests/" + g.snakeName + ".go"

	tmpl, err := template.New("request").Parse(requestTemplate)
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(g.prj.Path, []string{"requests"}, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return ErrRequestExists
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"camelName": g.camelName,
	}

	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	g.prj.Logger.Info("Request generated: " + g.camelName)

	return nil
}
