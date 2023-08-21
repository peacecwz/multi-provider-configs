package env

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	multiproviderconfigs "github.com/peacecwz/multi-provider-configs/providers"
)

type EnvConfigProvider[T any] struct {
	options *EnvConfigOptions
}

func (e EnvConfigProvider[T]) Parse(config *T) error {
	if err := env.Parse(config); err != nil {
		return fmt.Errorf("failed to parse env to struct. details: %+v", err)
	}

	return nil
}

func (e EnvConfigProvider[T]) IsRequired() bool {
	return e.options.Required
}

type EnvConfigOptions struct {
	Required bool
}

type EnvConfigOptionSetter func(*EnvConfigOptions)

func WithRequired(required bool) EnvConfigOptionSetter {
	return func(opt *EnvConfigOptions) {
		opt.Required = required
	}
}

func NewEnvConfigProvider[T any](opts ...EnvConfigOptionSetter) multiproviderconfigs.ConfigProvider[T] {
	options := &EnvConfigOptions{}
	for _, opt := range opts {
		opt(options)
	}

	return NewEnvConfigProviderWithOption[T](options)
}

func NewEnvConfigProviderWithOption[T any](opt *EnvConfigOptions) multiproviderconfigs.ConfigProvider[T] {
	return &EnvConfigProvider[T]{
		options: opt,
	}
}
