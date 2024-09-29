package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	"os/exec"
	"path"
	"text/template"
)

//go:embed sources/root/.air.toml.tmpl
var air_config_template []byte

func buildAirConfig(prj *project.Project) error {
	prj.Logger.Debug("Generating .air.toml")

	file, err := os.OpenFile(path.Join(prj.Path, ".air.toml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(air_config_template)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/template.gitignore
var gitignore_template []byte

func buildGitignore(prj *project.Project) error {
	prj.Logger.Debug("Generating .gitignore")

	file, err := os.OpenFile(path.Join(prj.Path, ".gitignore"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(gitignore_template)

	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/docker-compose.yml.tmpl
var docker_compose_template []byte

func buildDockerCompose(prj *project.Project) error {
	prj.Logger.Debug("Generating docker-compose.yml")

	tmpl, err := template.New("docker_compose").Parse(string(docker_compose_template))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "docker-compose.yml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"database":    prj.Database,
		"projectName": prj.Name,
	}

	return tmpl.Execute(file, data)
}

//go:embed sources/root/go.mod.tmpl
var go_mod_template []byte

func buildGoMod(prj *project.Project) error {
	prj.Logger.Debug("Generating go.mod")

	tmpl, err := template.New("go_mod").Parse(string(go_mod_template))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "go.mod"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Running go mod tidy")
	cmd := exec.Command("go", "mod", "tidy", "-v")
	cmd.Dir = prj.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/Taskfile.yml.tmpl
var taskfile_template []byte

func buildTaskfile(prj *project.Project) error {
	prj.Logger.Debug("Generating Taskfile.yml")

	tmpl, err := template.New("taskfile").Parse(string(taskfile_template))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(prj.Path, "Taskfile.yml"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	data := map[string]interface{}{
		"isInertia": prj.FrontendType.IsInertia(),
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func BuildRootFiles(prj *project.Project) error {
	err := buildAirConfig(prj)
	if err != nil {
		return err
	}

	err = buildGitignore(prj)
	if err != nil {
		return err
	}

	err = buildDockerCompose(prj)
	if err != nil {
		return err
	}

	err = buildGoMod(prj)
	if err != nil {
		return err
	}

	err = buildTaskfile(prj)
	if err != nil {
		return err
	}

	return nil
}
