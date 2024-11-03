package core

type RenovateConfig struct {
	Schema            string        `json:"$schema"`
	Extends           []string      `json:"extends"`
	PrHourlyLimit     int           `json:"prHourlyLimit"`
	PrConcurrentLimit int           `json:"prConcurrentLimit"`
	PackageRules      []PackageRule `json:"packageRules"`
	PlatformAutomerge bool          `json:"platformAutomerge"`
	PostUpdateOptions []string      `json:"postUpdateOptions,omitempty"`
	IgnorePaths       []string      `json:"ignorePaths,omitempty"`
}

type PackageRule struct {
	RangeStrategy     string   `json:"rangeStrategy,omitempty"`
	Automerge         bool     `json:"automerge,omitempty"`
	IgnoreTests       bool     `json:"ignoreTests,omitempty"`
	AutomergeType     string   `json:"automergeType,omitempty"`
	AutomergeStrategy string   `json:"automergeStrategy,omitempty"`
	MatchDepTypes     []string `json:"matchDepTypes,omitempty"`
	MatchUpdateTypes  []string `json:"matchUpdateTypes,omitempty"`
	RecreateWhen      string   `json:"recreateWhen,omitempty"`
	Enabled           bool     `json:"enabled,omitempty"`
	MatchManagers     []string `json:"matchManagers,omitempty"`
}
