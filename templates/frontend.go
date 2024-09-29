package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"os/exec"
	"path"
	"text/template"
)

func BuildFrontend(prj *project.Project) error {
	prj.Logger.Debug("Building frontend")
	if !prj.HasFrontend {
		prj.Logger.Debug("Project has no frontend, skipping")
		return nil
	}

	switch prj.FrontendType {
	case project.InertiaReact:
		err := buildReactInertiaFrontend(prj)
		if err != nil {
			return err
		}
	case project.InertiaVue:
		err := buildVueInertiaFrontend(prj)
		if err != nil {
			return err
		}
	case project.InertiaSvelte:
		err := buildSvelteInertiaFrontend(prj)
		if err != nil {
			return err
		}
	}

	err := buildFrontend_Frontend(prj)
	if err != nil {
		return err
	}

	err = installPackages(prj)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/frontend/frontend.go.tmpl
var frontend_template []byte

func buildFrontend_Frontend(prj *project.Project) error {
	prj.Logger.Debug("Generating frontend/frontend.go")

	tmpl, err := template.New("frontend_frontend").Parse(string(frontend_template))
	if err != nil {
		return err
	}

	err = helpers.CreateDirectories(prj.Path, []string{"frontend"}, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "frontend/frontend.go"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, nil)
}

func installPackages(prj *project.Project) error {
	prj.Logger.Debug("Installing packages")
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
