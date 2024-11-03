package core

import (
	"fmt"

	"github.com/gkwa/alpinerider/internal/types"
	"github.com/imdario/mergo"
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
