package core

import "github.com/gkwa/alpinerider/internal/types"

var BaseConfig = types.RenovateConfig{
	Schema: "https://docs.renovatebot.com/renovate-schema.json",
	Extends: []string{
		"config:best-practices",
		":dependencyDashboard",
	},
	PrHourlyLimit:     0,
	PrConcurrentLimit: 0,
	PlatformAutomerge: true,
}
