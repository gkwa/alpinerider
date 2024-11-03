package core

import (
	"github.com/gkwa/alpinerider/core/dog"
	"github.com/gkwa/alpinerider/core/types"
)

func (c *ConfigComposer) Dog() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = append(config.PackageRules, dog.Rule)
	return config, nil
}
