package core

import (
	"github.com/gkwa/alpinerider/core/cat"
	"github.com/gkwa/alpinerider/core/types"
)

func (c *ConfigComposer) Cat() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = append(config.PackageRules, cat.Rule)
	return config, nil
}
