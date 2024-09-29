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

func buildReactInertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building React Inertia frontend")

	cmd := exec.Command("bun", "create", "vite", "--template", "react-ts", "frontend")
	cmd.Dir = prj.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
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
var react_vite_config_template []byte

func buildViteReactConfig(prj *project.Project) error {
	prj.Logger.Debug("Building Vite React config at frontend/vite.config.ts")

	tmpl, err := template.New("react_vite_config").Parse(string(react_vite_config_template))
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

//go:embed sources/frontend/react/root.gohtml.tmpl
var react_root_template []byte

func buildReactRoot(prj *project.Project) error {
	prj.Logger.Debug("Building React root component at frontend/root.gohtml")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/root.gohtml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(react_root_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/react/package.json.tmpl
var react_package_json_template []byte

func buildReactPackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building React package.json")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/package.json"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(react_package_json_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/react/src
var react_src_template embed.FS

func buildReactSrc(prj *project.Project) error {
	prj.Logger.Debug("Building React src")

	//delete all files in src
	err := os.RemoveAll(path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend/src"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(react_src_template, "sources/frontend/react/src", path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	return nil
}
