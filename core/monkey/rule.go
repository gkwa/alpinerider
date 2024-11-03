package monkey

import "github.com/gkwa/alpinerider/core/types"

var Rule = types.PackageRule{
	Automerge:         true,
	IgnoreTests:       true,
	AutomergeType:     "branch",
	AutomergeStrategy: "merge-commit",
	MatchDepTypes: []string{
		"*",
	},
	MatchUpdateTypes: []string{
		"minor",
		"patch",
		"pin",
		"digest",
		"replacement",
	},
}