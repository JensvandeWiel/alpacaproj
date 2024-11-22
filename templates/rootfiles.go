package templates

import (
	_ "embed"
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

//go:embed sources/root/.air.toml.tmpl
var airConfigTemplate string

func buildAirConfig(prj *project.Project) error {
	prj.Logger.Debug("Generating .air.toml")

	data := map[string]interface{}{
		"frontendType": prj.FrontendType,
	}

	err := helpers.WriteTemplateToFile(prj, ".air.toml", airConfigTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/template.gitignore
var gitignoreTemplate string

func buildGitignore(prj *project.Project) error {
	prj.Logger.Debug("Generating .gitignore")

	err := helpers.WriteTemplateToFile(prj, ".gitignore", gitignoreTemplate, nil)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/docker-compose.yml.tmpl
var dockerComposeTemplate string

func buildDockerCompose(prj *project.Project) error {
	prj.Logger.Debug("Generating docker-compose.yml")

	data := map[string]interface{}{
		"database":    prj.Database,
		"projectName": prj.Name,
	}

	err := helpers.WriteTemplateToFile(prj, "docker-compose.yml", dockerComposeTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/go.mod.tmpl
var goModTemplate string

func buildGoMod(prj *project.Project) error {
	prj.Logger.Debug("Generating go.mod")

	data := map[string]interface{}{
		"packageName": prj.PackageName,
	}

	err := helpers.WriteTemplateToFile(prj, "go.mod", goModTemplate, data)
	if err != nil {
		return err
	}

	return nil
}

//go:embed sources/root/Taskfile.yml.tmpl
var taskfileTemplate string

func buildTaskfile(prj *project.Project) error {
	prj.Logger.Debug("Generating Taskfile.yml")

	data := map[string]interface{}{
		"isInertia": prj.FrontendType.IsInertia(),
	}

	err := helpers.WriteTemplateToFile(prj, "Taskfile.yml", taskfileTemplate, data)
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
