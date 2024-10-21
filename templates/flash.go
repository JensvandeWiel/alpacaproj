package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/flash/flash.go.tmpl
var flashTemplate string

// BuildFlash_Flash generates the flash/flash.go file
func BuildFlash_Flash(prj *project.Project) error {
	prj.Logger.Debug("Generating flash/flash.go")
	if !prj.FrontendType.IsInertia() {
		prj.Logger.Debug("Frontend type is not inertia, skipping")
		return nil
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "flash/flash.go", flashTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

func BuildFlash(prj *project.Project) error {
	err := BuildFlash_Flash(prj)
	if err != nil {
		return err
	}

	return nil
}
