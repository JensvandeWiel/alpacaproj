package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
)

//go:embed sources/helpers/is_dev.go.tmpl
var is_dev_template []byte

//go:embed sources/helpers/is_release.go.tmpl
var is_prod_template []byte

// buildHelpers_IsDev generates the helpers/is_dev.go file
func buildHelpers_IsDev(prj *project.Project) error {
	prj.Logger.Debug("Generating helpers/is_dev.go")

	tmpl, err := template.New("is_dev").Parse(string(is_dev_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"helpers"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "helpers/is_dev.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	return tmpl.Execute(file, nil)
}

// buildHelpers_IsProd generates the helpers/is_prod.go file

func buildHelpers_IsProd(prj *project.Project) error {
	prj.Logger.Debug("Generating helpers/is_prod.go")

	tmpl, err := template.New("is_prod").Parse(string(is_prod_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"helpers"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "helpers/is_prod.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	return tmpl.Execute(file, nil)
}

func BuildHelpers(prj *project.Project) error {
	err := buildHelpers_IsDev(prj)
	if err != nil {
		return err
	}

	err = buildHelpers_IsProd(prj)
	if err != nil {
		return err
	}

	return nil
}
