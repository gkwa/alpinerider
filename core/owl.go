package core

import (
	"github.com/gkwa/alpinerider/core/owl"
	"github.com/gkwa/alpinerider/internal/types"
)

func (c *ConfigComposer) Owl() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = []types.PackageRule{owl.Rule}
	config.PostUpdateOptions = []string{
		"gomodTidy",
		"gomodUpdateImportPaths",
	}
	config.IgnorePaths = []string{
		"**/testdata/go.mod",
	}
	return config, nil
}
