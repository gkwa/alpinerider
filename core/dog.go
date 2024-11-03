package core

import (
	"github.com/gkwa/alpinerider/core/dog"
	"github.com/gkwa/alpinerider/internal/types"
)

func (c *ConfigComposer) Dog() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = []types.PackageRule{dog.Rule}
	return config, nil
}
