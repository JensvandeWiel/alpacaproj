package extras

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"os/exec"
	"path"
)

func BuildSvelte5(prj *project.Project) error {
	prj.Logger.Debug("Building extra Svelte 5 frontend")

	err := buildPackageJSON(prj)
	if err != nil {
		return err
	}

	err = buildMainTS(prj)
	if err != nil {
		return err
	}

	err = installPackages(prj)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/svelte5/package.json.tmpl
var svelte5_package_json_template []byte

func buildPackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte 5 package.json")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/package.json"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Truncate(path.Join(prj.Path, "frontend/package.json"), 0)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(svelte5_package_json_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed templates/svelte5/main.ts.tmpl
var svelte5_main_ts_template []byte

func buildMainTS(prj *project.Project) error {
	prj.Logger.Debug("Building Svelte 5 main.ts")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/src/main.ts"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Truncate(path.Join(prj.Path, "frontend/src/main.ts"), 0)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(svelte5_main_ts_template)

	if err != nil {
		return err
	}

	return nil
}

func installPackages(prj *project.Project) error {
	prj.Logger.Debug("Installing Svelte 5 packages")

	cmd := exec.Command("bun", "install")
	cmd.Dir = path.Join(prj.Path, "frontend")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
