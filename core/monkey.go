package core

import (
	"github.com/gkwa/alpinerider/core/monkey"
	"github.com/gkwa/alpinerider/core/types"
)

func (c *ConfigComposer) Monkey() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = append(config.PackageRules, monkey.Rule)
	return config, nil
}
