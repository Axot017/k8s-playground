module github.com/Axot017/k8s-playground/gateway

go 1.21.3

require (
	github.com/Axot017/k8s-playground/common v0.0.0
	github.com/labstack/echo/v4 v4.11.3
	github.com/segmentio/ksuid v1.0.4
	go.uber.org/fx v1.20.1
)

replace github.com/Axot017/k8s-playground/common => ../common

require (
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
