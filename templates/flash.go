package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"path"
	"text/template"
)

//go:embed sources/flash/flash.go.tmpl
var flash_template []byte

// BuildFlash_Flash generates the flash/flash.go file
func BuildFlash_Flash(prj *project.Project) error {
	prj.Logger.Debug("Generating flash/flash.go")
	if !prj.FrontendType.IsInertia() {
		prj.Logger.Debug("Frontend type is not inertia, skipping")
		return nil
	}

	tmpl, err := template.New("flash_flash").Parse(string(flash_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"flash"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "flash/flash.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	return tmpl.Execute(file, data)
}

func BuildFlash(prj *project.Project) error {
	err := BuildFlash_Flash(prj)
	if err != nil {
		return err
	}

	return nil
}
