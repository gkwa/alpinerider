package core

import (
	"fmt"

	"dario.cat/mergo"
	"github.com/gkwa/alpinerider/core/types"
)

type ConfigComposer struct {
	baseConfig types.RenovateConfig
}

func NewConfigComposer(base types.RenovateConfig) *ConfigComposer {
	return &ConfigComposer{
		baseConfig: base,
	}
}

func MergeConfigs(configs ...types.RenovateConfig) (types.RenovateConfig, error) {
	result := types.RenovateConfig{}
	for _, config := range configs {
		if err := mergo.Merge(&result, config, mergo.WithOverride); err != nil {
			return types.RenovateConfig{}, fmt.Errorf("failed to merge configs: %w", err)
		}
	}
	return result, nil
}
