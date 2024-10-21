package extras

import (
	"github.com/JensvandeWiel/alpacaproj/helpers"
	"github.com/JensvandeWiel/alpacaproj/project"
	"sort"
)

// Define the priority map
var extraPriority = map[project.ExtraOption]int{
	project.Svelte5:      0,
	project.SQLC:         0,
	project.FrontendAuth: 1,
}

func ApplyExtras(prj *project.Project) error {
	// Sort extras based on the priority map
	sort.Slice(prj.Extras, func(i, j int) bool {
		return extraPriority[prj.Extras[i]] < extraPriority[prj.Extras[j]]
	})

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
		case project.FrontendAuth:
			err := BuildFrontendAuth(prj)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
