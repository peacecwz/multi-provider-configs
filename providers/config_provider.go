package providers

type ConfigProvider[T any] interface {
	Parse(config *T) error
	IsRequired() bool
}
