package templates

import (
	"embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
)

func buildReactInertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building React Inertia frontend")

	err := helpers.RunCommand(prj.Path, true, "bun", "create", "vite", "--template", "react-ts", "frontend")
	if err != nil {
		return err
	}

	err = buildViteReactConfig(prj)
	if err != nil {
		return err
	}
	err = buildReactRoot(prj)
	if err != nil {
		return err
	}

	err = buildReactPackageJSON(prj)
	if err != nil {
		return err
	}

	err = buildReactSrc(prj)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/react/vite.config.ts.tmpl
var reactViteConfigTemplate string

func buildViteReactConfig(prj *project.Project) error {
	prj.Logger.Debug("Building Vite React config at frontend/vite.config.ts")

	err := helpers.WriteTemplateToFile(prj, "frontend/vite.config.ts", reactViteConfigTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/react/root.gohtml.tmpl
var reactRootTemplate string

func buildReactRoot(prj *project.Project) error {
	prj.Logger.Debug("Building React root component at frontend/root.gohtml")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/root.gohtml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(reactRootTemplate))
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/react/package.json.tmpl
var reactPackageJSONTemplate string

func buildReactPackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building React package.json")

	err := helpers.WriteTemplateToFile(prj, "frontend/package.json", reactPackageJSONTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/react/src
var reactSrcTemplate embed.FS

func buildReactSrc(prj *project.Project) error {
	prj.Logger.Debug("Building React src")

	// delete all files in src
	err := os.RemoveAll(path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend/src"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(reactSrcTemplate, "sources/frontend/react/src", path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	return nil
}
