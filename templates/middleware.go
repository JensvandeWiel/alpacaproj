package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/middleware/helpers.go.tmpl
var middlewareHelpersTemplate string

//go:embed sources/middleware/middleware.go.tmpl
var middlewareTemplate string

// buildMiddleware_Middleware generates the middleware/middleware.go file
func buildMiddleware_Middleware(prj *project.Project) error {
	prj.Logger.Debug("Generating middleware/middleware.go")

	err := helpers.CreateDirectories(prj.Path, []string{"middleware"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.WriteTemplateToFile(prj, "middleware/middleware.go", middlewareTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

// buildMiddleware_Helpers generates the middleware/helpers.go file
func buildMiddleware_Helpers(prj *project.Project) error {
	prj.Logger.Debug("Generating middleware/helpers.go")

	err := helpers.CreateDirectories(prj.Path, []string{"middleware"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.WriteTemplateToFile(prj, "middleware/helpers.go", middlewareHelpersTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

// BuildMiddleware generates the middleware files
func BuildMiddleware(prj *project.Project) error {
	err := buildMiddleware_Middleware(prj)
	if err != nil {
		return err
	}

	err = buildMiddleware_Helpers(prj)
	if err != nil {
		return err
	}

	return nil
}
