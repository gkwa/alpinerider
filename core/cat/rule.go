package cat

import (
	"github.com/gkwa/alpinerider/core/common"
	"github.com/gkwa/alpinerider/internal/types"
)

var Rule = types.PackageRule{
	RangeStrategy:     "pin",
	Automerge:         true,
	IgnoreTests:       true,
	AutomergeType:     "branch",
	AutomergeStrategy: "merge-commit",
	MatchDepTypes: []string{
		"*",
	},
	MatchUpdateTypes: common.MatchUpdateTypes,
}
