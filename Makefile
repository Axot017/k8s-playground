gateway:
	go run gateway/cmd/service/main.go

profile:
	go run profile/cmd/service/main.go

minikube:
	minikube start --kubernetes-version=v1.23.0 --memory=6g --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0

chart:
	docker build ./gateway -t gateway-local -f ./deployments/Dockerfile
	docker build ./profile -t profile-local -f ./deployments/Dockerfile
	helm install local-app ./deployments/chart

uninstall:
	helm uninstall local-app

local:
	docker-compose -f deployments/docker-compose.yml up --build -d
