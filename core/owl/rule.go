package owl

import "github.com/gkwa/alpinerider/core/types"

var Rule = types.PackageRule{
	Automerge:    true,
	RecreateWhen: "always",
	MatchUpdateTypes: []string{
		"minor",
		"patch",
		"pin",
		"digest",
		"replacement",
	},
}
