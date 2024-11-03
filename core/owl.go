package core

import (
	"github.com/gkwa/alpinerider/core/owl"
	"github.com/gkwa/alpinerider/core/types"
)

func (c *ConfigComposer) Owl() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = append(config.PackageRules, owl.Rule)
	config.IgnorePaths = []string{
		"**/testdata/go.mod",
	}
	return config, nil
}
