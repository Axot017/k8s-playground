module github.com/Axot017/k8s-playground/gateway

go 1.21.3

require (
	github.com/Axot017/k8s-playground/common v0.0.0
	github.com/go-chi/chi/v5 v5.0.10
	go.uber.org/fx v1.20.1
)

replace github.com/Axot017/k8s-playground/common => ../common

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)
