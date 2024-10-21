package templates

import (
	"embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
)

func buildSvelteInertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte Inertia frontend")

	err := helpers.RunCommand(prj.Path, true, "bun", "create", "vite", "--template", "svelte-ts", "frontend")
	if err != nil {
		return err
	}

	err = buildViteSvelteConfig(prj)
	if err != nil {
		return err
	}
	err = buildSvelteRoot(prj)
	if err != nil {
		return err
	}

	err = buildSveltePackageJSON(prj)
	if err != nil {
		return err
	}

	err = buildSvelteSrc(prj)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/svelte/vite.config.ts.tmpl
var svelteViteConfigTemplate string

func buildViteSvelteConfig(prj *project.Project) error {
	prj.Logger.Debug("Building Vite Svelte config at frontend/vite.config.ts")

	err := helpers.WriteTemplateToFile(prj, "frontend/vite.config.ts", svelteViteConfigTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/svelte/root.gohtml.tmpl
var svelteRootTemplate string

func buildSvelteRoot(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte root component at frontend/root.gohtml")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/root.gohtml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(svelteRootTemplate))
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/svelte/package.json.tmpl
var sveltePackageJSONTemplate string

func buildSveltePackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte package.json")

	err := helpers.WriteTemplateToFile(prj, "frontend/package.json", sveltePackageJSONTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/svelte/src
var svelteSrcTemplate embed.FS

func buildSvelteSrc(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte src")

	// delete all files in src
	err := os.RemoveAll(path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend/src"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(svelteSrcTemplate, "sources/frontend/svelte/src", path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	return nil
}
