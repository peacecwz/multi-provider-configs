package json

import (
	"encoding/json"
	"fmt"
	multiproviderconfigs "github.com/peacecwz/multi-provider-configs/providers"
	"os"
)

type JsonConfigProvider[T any] struct {
	options *JsonConfigOptions
}

func (j JsonConfigProvider[T]) Parse(config *T) error {
	configData, err := os.ReadFile(j.options.ConfigFile)
	if err != nil {
		if !j.options.Required {
			return nil
		}

		return fmt.Errorf("cannot read %s file. details: %+v", j.options.ConfigFile, err)
	}

	if err := json.Unmarshal(configData, config); err != nil && j.options.Required {
		return fmt.Errorf("failed to unmarshall config. details: %+v", err)
	}

	return nil
}

func (j JsonConfigProvider[T]) IsRequired() bool {
	return j.options.Required
}

type JsonConfigOptions struct {
	Required   bool
	ConfigFile string
}

type JsonConfigOptionSetter func(*JsonConfigOptions)

func WithRequired(required bool) JsonConfigOptionSetter {
	return func(opt *JsonConfigOptions) {
		opt.Required = required
	}
}

func WithConfigFile(configFile string) JsonConfigOptionSetter {
	return func(opt *JsonConfigOptions) {
		opt.ConfigFile = configFile
	}
}

func NewJsonConfigProvider[T any](opts ...JsonConfigOptionSetter) multiproviderconfigs.ConfigProvider[T] {
	options := &JsonConfigOptions{}
	for _, opt := range opts {
		opt(options)
	}

	return NewJsonConfigProviderWithOption[T](options)
}

func NewJsonConfigProviderWithOption[T any](opt *JsonConfigOptions) multiproviderconfigs.ConfigProvider[T] {
	return &JsonConfigProvider[T]{
		options: opt,
	}
}
