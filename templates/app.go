package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
)

//go:embed sources/app/app.go.tmpl
var app_template []byte

// buildApp_App generates the app/app.go file
func buildApp_App(prj *project.Project) error {
	prj.Logger.Debug("Generating app/app.go")
	tmpl, err := template.New("frontend_inertia").Parse(string(app_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"app"}, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "app/app.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]interface{}{
		"hasFrontend": prj.HasFrontend,
		"isInertia":   prj.FrontendType.IsInertia(),
	}

	return tmpl.Execute(file, data)
}

//go:embed sources/app/inertia_frontend.go.tmpl
var inertia_frontend_template []byte

// buildApp_InertiaFrontend generates the app/inertia_frontend.go file
func buildApp_InertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Generating app/inertia_frontend.go")
	if !prj.FrontendType.IsInertia() {
		prj.Logger.Debug("Frontend type is not inertia, skipping")
		return nil
	}

	tmpl, err := template.New("inertia_frontend").Parse(string(inertia_frontend_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"app"}, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "app/inertia_frontend.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
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
