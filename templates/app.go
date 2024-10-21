package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/app/app.go.tmpl
var appTemplate string

// buildApp_App generates the app/app.go file
func buildApp_App(prj *project.Project) error {
	prj.Logger.Debug("Generating app/app.go")

	data := map[string]interface{}{
		"hasFrontend": prj.HasFrontend,
		"isInertia":   prj.FrontendType.IsInertia(),
	}

	err := helpers.WriteTemplateToFile(prj, "app/app.go", appTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/app/inertia_frontend.go.tmpl
var inertiaFrontendTemplate string

// buildApp_InertiaFrontend generates the app/inertia_frontend.go file
func buildApp_InertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Generating app/inertia_frontend.go")
	if !prj.FrontendType.IsInertia() {
		prj.Logger.Debug("Frontend type is not inertia, skipping")
		return nil
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "app/inertia_frontend.go", inertiaFrontendTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

func BuildApp(prj *project.Project) error {
	err := buildApp_App(prj)
	if err != nil {
		return err
	}
	err = buildApp_InertiaFrontend(prj)
	if err != nil {
		return err
	}
	return nil
}
