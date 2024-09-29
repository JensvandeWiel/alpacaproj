package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
)

//go:embed sources/middleware/helpers.go.tmpl
var middleware_helpers_template []byte

//go:embed sources/middleware/middleware.go.tmpl
var middleware_template []byte

// buildMiddleware_Middleware generates the middleware/middleware.go file
func buildMiddleware_Middleware(prj *project.Project) error {
	prj.Logger.Debug("Generating middleware/middleware.go")

	err := helpers.CreateDirectories(prj.Path, []string{"middleware"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "middleware/middleware.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(middleware_template)
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

	file, err := os.OpenFile(path.Join(prj.Path, "middleware/helpers.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(middleware_helpers_template)
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
