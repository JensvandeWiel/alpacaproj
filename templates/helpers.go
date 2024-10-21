package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/helpers/is_dev.go.tmpl
var isDevTemplate string

//go:embed sources/helpers/is_release.go.tmpl
var isProdTemplate string

// buildHelpers_IsDev generates the helpers/is_dev.go file
func buildHelpers_IsDev(prj *project.Project) error {
	prj.Logger.Debug("Generating helpers/is_dev.go")

	err := helpers.WriteTemplateToFile(prj, "helpers/is_dev.go", isDevTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

// buildHelpers_IsProd generates the helpers/is_prod.go file
func buildHelpers_IsProd(prj *project.Project) error {
	prj.Logger.Debug("Generating helpers/is_prod.go")

	err := helpers.WriteTemplateToFile(prj, "helpers/is_prod.go", isProdTemplate, nil)
	if err != nil {
		return err
	}

	return nil
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
