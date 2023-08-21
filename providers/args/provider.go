package args

import (
	"github.com/alexflint/go-arg"
	multiproviderconfigs "github.com/peacecwz/multi-provider-configs/providers"
)

type ArgsConfigProvider[T any] struct {
	options *ArgsConfigOptions
}

func (p ArgsConfigProvider[T]) Parse(config *T) error {
	_ = arg.Parse(config)

	return nil
}

func (p ArgsConfigProvider[T]) IsRequired() bool {
	return p.options.Required
}

type ArgsConfigOptions struct {
	Required bool
}

type ArgsConfigOptionSetter func(*ArgsConfigOptions)

func WithRequired(required bool) ArgsConfigOptionSetter {
	return func(opt *ArgsConfigOptions) {
		opt.Required = required
	}
}

func NewArgsConfigProvider[T any](opts ...ArgsConfigOptionSetter) multiproviderconfigs.ConfigProvider[T] {
	options := &ArgsConfigOptions{}
	for _, opt := range opts {
		opt(options)
	}

	return NewArgsConfigProviderWithOption[T](options)
}

func NewArgsConfigProviderWithOption[T any](opt *ArgsConfigOptions) multiproviderconfigs.ConfigProvider[T] {
	return &ArgsConfigProvider[T]{
		options: opt,
	}
}
