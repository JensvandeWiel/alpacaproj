package extras

import (
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
)

func ApplyExtras(prj *project.Project) error {
	for _, extra := range prj.Extras {
		switch extra {
		case project.Svelte5:
			if prj.FrontendType != project.InertiaSvelte {
				prj.Logger.Warn("Svelte 5 is only available for Inertia+Svelte frontend")
				continue
			}
			err := BuildSvelte5(prj)
			if err != nil {
				return err
			}
		case project.SQLC:
			err := helpers.CreateDirectories(prj.Path, []string{"queries", "repository"}, 0755)
			if err != nil {
				return err
			}

			err = BuildSQLC(prj)
			if err != nil {
				return err
			}
			continue
		}
	}

	return nil
}
