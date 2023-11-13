gateway:
	go run gateway/cmd/service/main.go

profile:
	go run profile/cmd/service/main.go

chart:
	docker build ./gateway -t gateway-local -f ./deployments/Dockerfile
	docker build ./profile -t profile-local -f ./deployments/Dockerfile
	helm install local-app ./deployments/chart

uninstall:
	helm uninstall local-app

local:
	docker-compose -f deployments/docker-compose.yml up --build -d
