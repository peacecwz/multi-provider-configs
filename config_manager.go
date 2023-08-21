package multi_provider_configs

import (
	"github.com/peacecwz/multi-provider-configs/providers"
	"github.com/peacecwz/multi-provider-configs/providers/args"
	"github.com/peacecwz/multi-provider-configs/providers/env"
	"github.com/peacecwz/multi-provider-configs/providers/json"
)

type ConfigManager[T any] struct {
	providers *Queue[providers.ConfigProvider[T]]
	config    T
}

func New[T any](value T) *ConfigManager[T] {
	configProviders := NewQueue[providers.ConfigProvider[T]]()

	configProviders.Enqueue(args.NewArgsConfigProvider[T](args.WithRequired(true)))
	configProviders.Enqueue(env.NewEnvConfigProvider[T](env.WithRequired(true)))
	configProviders.Enqueue(json.NewJsonConfigProvider[T](json.WithRequired(false), json.WithConfigFile("config.json")))

	return &ConfigManager[T]{
		providers: configProviders,
		config:    value,
	}
}

func (m *ConfigManager[T]) Load() (*T, error) {
	var err error
	for !m.providers.IsEmpty() {
		provider, ok := m.providers.Dequeue()
		if !ok {
			continue
		}

		err = provider.Parse(&m.config)
		if err != nil && provider.IsRequired() {
			panic(err)
		}
	}

	return &m.config, err
}
