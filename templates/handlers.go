package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
)

//go:embed sources/handlers/handler.go.tmpl
var handler_template []byte

func buildHandlers_Handler(prj *project.Project) error {
	prj.Logger.Debug("Generating handlers/handler.go")

	tmpl, err := template.New("handler").Parse(string(handler_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"handlers"}, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "handlers/handler.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
}

//go:embed sources/handlers/api_handler.go.tmpl
var api_handler_template []byte

func buildHandlers_APIHandler(prj *project.Project) error {
	prj.Logger.Debug("Generating handlers/api_handler.go")

	tmpl, err := template.New("api_handler").Parse(string(api_handler_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"handlers"}, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "handlers/api_handler.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
}

//go:embed sources/handlers/main_handler.go.tmpl
var main_handler_template []byte

func buildHandlers_MainHandler(prj *project.Project) error {
	prj.Logger.Debug("Generating handlers/main_handler.go")

	if !prj.FrontendType.IsInertia() {
		prj.Logger.Debug("Frontend type is not inertia, skipping")
		return nil
	}

	tmpl, err := template.New("main_handler").Parse(string(main_handler_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"handlers"}, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "handlers/main_handler.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
}

func BuildHandlers(prj *project.Project) error {
	err := buildHandlers_Handler(prj)
	if err != nil {
		return err
	}

	err = buildHandlers_APIHandler(prj)
	if err != nil {
		return err
	}

	err = buildHandlers_MainHandler(prj)
	if err != nil {
		return err
	}
	return nil
}
