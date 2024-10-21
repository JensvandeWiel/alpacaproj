package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/handlers/handler.go.tmpl
var handlerTemplate string

// buildHandlers_Handler generates the handlers/handler.go file
func buildHandlers_Handler(prj *project.Project) error {
	prj.Logger.Debug("Generating handlers/handler.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "handlers/handler.go", handlerTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/handlers/api_handler.go.tmpl
var apiHandlerTemplate string

// buildHandlers_APIHandler generates the handlers/api_handler.go file
func buildHandlers_APIHandler(prj *project.Project) error {
	prj.Logger.Debug("Generating handlers/api_handler.go")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "handlers/api_handler.go", apiHandlerTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/handlers/main_handler.go.tmpl
var mainHandlerTemplate string

// buildHandlers_MainHandler generates the handlers/main_handler.go file
func buildHandlers_MainHandler(prj *project.Project) error {
	prj.Logger.Debug("Generating handlers/main_handler.go")

	if !prj.FrontendType.IsInertia() {
		prj.Logger.Debug("Frontend type is not inertia, skipping")
		return nil
	}

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "handlers/main_handler.go", mainHandlerTemplate, data)
	if err != nil {
		return err
	}

	return nil
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
