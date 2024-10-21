package helpers

import (
	"github.com/JensvandeWiel/alpacaproj/project"
	path2 "path"
)

// InstallNPMPackages installs NPM packages in the given path. The path is relative to the project path.
func InstallNPMPackages(prj *project.Project, path string) error {
	prj.Logger.Debug("Installing NPM packages")

	err := RunCommand(path2.Join(prj.Path, path), true, "bun", "install")
	if err != nil {
		return err
	}

	prj.Logger.Debug("Installed NPM packages")

	return nil
}
