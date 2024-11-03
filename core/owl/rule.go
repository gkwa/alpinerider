package owl

import (
	"github.com/gkwa/alpinerider/core/common"
	"github.com/gkwa/alpinerider/internal/types"
)

var Rule = types.PackageRule{
	Automerge:        true,
	RecreateWhen:     "always",
	MatchUpdateTypes: append(common.MatchUpdateTypes, "replacement"),
}
