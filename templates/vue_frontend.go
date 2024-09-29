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

func buildVueInertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building Vue Inertia frontend")

	cmd := exec.Command("bun", "create", "vite", "--template", "vue-ts", "frontend")
	cmd.Dir = prj.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}

	err = buildViteVueConfig(prj)
	if err != nil {
		return err
	}
	err = buildVueRoot(prj)
	if err != nil {
		return err
	}

	err = buildVuePackageJSON(prj)
	if err != nil {
		return err
	}

	err = buildVueSrc(prj)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/vue/vite.config.ts.tmpl
var vue_vite_config_template []byte

func buildViteVueConfig(prj *project.Project) error {
	prj.Logger.Debug("Building Vite Vue config at frontend/vite.config.ts")

	tmpl, err := template.New("vue_vite_config").Parse(string(vue_vite_config_template))
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

//go:embed sources/frontend/vue/root.gohtml.tmpl
var vue_root_template []byte

func buildVueRoot(prj *project.Project) error {
	prj.Logger.Debug("Building Vue root component at frontend/root.gohtml")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/root.gohtml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(vue_root_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/vue/package.json.tmpl
var vue_package_json_template []byte

func buildVuePackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building Vue package.json")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/package.json"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(vue_package_json_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/vue/src
var vue_src_template embed.FS

func buildVueSrc(prj *project.Project) error {
	prj.Logger.Debug("Building Vue src")

	//delete all files in src
	err := os.RemoveAll(path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend/src"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(vue_src_template, "sources/frontend/vue/src", path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	return nil
}
