package dog

import (
	"github.com/gkwa/alpinerider/core/common"
	"github.com/gkwa/alpinerider/internal/types"
)

var Rule = types.PackageRule{
	Automerge:        true,
	RecreateWhen:     "always",
	MatchUpdateTypes: common.MatchUpdateTypes,
}
