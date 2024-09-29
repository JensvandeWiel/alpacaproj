package templates

import (
	"embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"os/exec"
	"path"
	"text/template"
)

func buildSvelteInertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte Inertia frontend")

	cmd := exec.Command("bun", "create", "vite", "--template", "svelte-ts", "frontend")
	cmd.Dir = prj.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
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
var svelte_vite_config_template []byte

func buildViteSvelteConfig(prj *project.Project) error {
	prj.Logger.Debug("Building Vite Svelte config at frontend/vite.config.ts")

	tmpl, err := template.New("svelte_vite_config").Parse(string(svelte_vite_config_template))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/vite.config.ts"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	return tmpl.Execute(file, nil)
}

//go:embed sources/frontend/svelte/root.gohtml.tmpl
var svelte_root_template []byte

func buildSvelteRoot(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte root component at frontend/root.gohtml")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/root.gohtml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(svelte_root_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/svelte/package.json.tmpl
var svelte_package_json_template []byte

func buildSveltePackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte package.json")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/package.json"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(svelte_package_json_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/svelte/src
var svelte_src_template embed.FS

func buildSvelteSrc(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte src")

	//delete all files in src
	err := os.RemoveAll(path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend/src"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(svelte_src_template, "sources/frontend/svelte/src", path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	return nil
}
