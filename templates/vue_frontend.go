package templates

import (
	"embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
)

func buildVueInertiaFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building Vue Inertia frontend")

	err := helpers.RunCommand(prj.Path, true, "bun", "create", "vite", "--template", "vue-ts", "frontend")
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
var vueViteConfigTemplate string

func buildViteVueConfig(prj *project.Project) error {
	prj.Logger.Debug("Building Vite Vue config at frontend/vite.config.ts")

	err := helpers.WriteTemplateToFile(prj, "frontend/vite.config.ts", vueViteConfigTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/vue/root.gohtml.tmpl
var vueRootTemplate string

func buildVueRoot(prj *project.Project) error {
	prj.Logger.Debug("Building Vue root component at frontend/root.gohtml")

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/root.gohtml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(vueRootTemplate))
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/vue/package.json.tmpl
var vuePackageJSONTemplate string

func buildVuePackageJSON(prj *project.Project) error {
	prj.Logger.Debug("Building Vue package.json")

	err := helpers.WriteTemplateToFile(prj, "frontend/package.json", vuePackageJSONTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/vue/src
var vueSrcTemplate embed.FS

func buildVueSrc(prj *project.Project) error {
	prj.Logger.Debug("Building Vue src")

	// delete all files in src
	err := os.RemoveAll(path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend/src"}, 0755)
	if err != nil {
		return err
	}

	err = helpers.CopyEmbeddedFiles(vueSrcTemplate, "sources/frontend/vue/src", path.Join(prj.Path, "frontend/src"))
	if err != nil {
		return err
	}

	return nil
}
