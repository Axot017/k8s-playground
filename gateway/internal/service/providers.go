package service

func Providers() []interface{} {
	return []interface{}{
		NewHttp,
	}
}
