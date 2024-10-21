package extras

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

func BuildSvelte5(prj *project.Project) error {
	prj.Logger.Debug("Building extra Svelte 5 frontend")

	err := buildPackageJSON(prj)
	if err != nil {
		return err
	}

	err = buildMainTS(prj)
	if err != nil {
		return err
	}

	err = helpers.InstallNPMPackages(prj, "frontend")
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/svelte5/package.json.tmpl
var svelte5PackageJSONTemplate string

func buildPackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte 5 package.json")

	err := helpers.WriteTemplateToFile(prj, "frontend/package.json", svelte5PackageJSONTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/svelte5/main.ts.tmpl
var svelte5MainTSTemplate string

func buildMainTS(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte 5 main.ts")

	err := helpers.WriteTemplateToFile(prj, "frontend/src/main.ts", svelte5MainTSTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}
