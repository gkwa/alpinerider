package core

import (
	"github.com/gkwa/alpinerider/core/monkey"
	"github.com/gkwa/alpinerider/internal/types"
)

func (c *ConfigComposer) Monkey() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = []types.PackageRule{monkey.Rule}
	config.PostUpdateOptions = []string{
		"gomodTidyE",
		"gomodMassage",
		"gomodUpdateImportPaths",
	}
	return config, nil
}
