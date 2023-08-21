package multi_provider_configs

import (
	"testing"
)

type Config struct {
	Test string `json:"test" env:"test" arg:"--test"`
}

func TestNew(t *testing.T) {
	// given
	var config Config
	configurationManager := New[Config](config)

	// when
	_, err := configurationManager.Load()

	// then
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if configurationManager.config.Test != "hello" {
		t.Errorf("unexpected config value: %v", configurationManager.config.Test)
	}
}
