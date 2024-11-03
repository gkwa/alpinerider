package core

import "github.com/gkwa/alpinerider/core/types"

var BaseConfig = types.RenovateConfig{
	Schema: "https://docs.renovatebot.com/renovate-schema.json",
	Extends: []string{
		"config:best-practices",
		":dependencyDashboard",
	},
	PrHourlyLimit:     0,
	PrConcurrentLimit: 0,
	PlatformAutomerge: true,
	PostUpdateOptions: []string{
		"gomodTidyE",
		"gomodMassage",
		"gomodUpdateImportPaths",
	},
	PackageRules: []types.PackageRule{
		{
			MatchManagers: []string{"gomod"},
			MatchDepTypes: []string{"indirect"},
			Enabled:       true,
		},
	},
}
