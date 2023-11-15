gateway:
	go run gateway/cmd/service/main.go

profile:
	go run profile/cmd/service/main.go

minikube:
	minikube start --kubernetes-version=v1.23.0 --memory=2g --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0

chart:
	docker build ./gateway -t gateway-local -f ./deployments/Dockerfile
	docker build ./profile -t profile-local -f ./deployments/Dockerfile
	helm install plg-monitoring ./deployments/monitoring-chart -n monitoring
	helm install local-app ./deployments/app-chart

uninstall:
	helm uninstall local-app
	helm uninstall plg-monitoring -n monitoring

local:
	docker-compose -f deployments/docker-compose.yml up --build -d
