package helpers

import (
	"github.com/JensvandeWiel/alpacaproj/project"
	"os"
	path2 "path"
	"text/template"
)

// WriteTemplateToFile writes a template to a file on the given path. PATH should not include the project path.
func WriteTemplateToFile(prj *project.Project, path string, tmplString string, data interface{}) error {
	filename := path2.Base(path)
	err := CreateDirectories(prj.Path, []string{
		path2.Dir(path),
	}, 0755)
	if err != nil {
		return err
	}

	prj.Logger.Debug("Writing template to file", "path", path)

	tmpl, err := template.New(filename).Parse(tmplString)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path2.Join(prj.Path, path), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	//fully empty the file
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
