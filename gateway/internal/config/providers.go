package config

func Providers() []interface{} {
	return []interface{}{
		NewConfig,
	}
}
