package core

import (
	"github.com/gkwa/alpinerider/core/cat"
	"github.com/gkwa/alpinerider/internal/types"
)

func (c *ConfigComposer) Cat() (types.RenovateConfig, error) {
	config := c.baseConfig
	config.PackageRules = []types.PackageRule{cat.Rule}
	return config, nil
}
