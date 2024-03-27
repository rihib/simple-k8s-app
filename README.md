# Simple k8s App

## Setup

```zsh
brew install minikube  # See https://minikube.sigs.k8s.io/docs/start/
minikube start  # Start minikube
minikube kubectl -- get po -A  # Get all pods
alias kubectl="minikube kubectl --"  # Set kubectl alias
minikube addons enable metrics-server  # Enable metrics-server for dashboard
minikube dashboard  # Open dashboard
```

## Deployment

See [this](https://www.baeldung.com/ops/docker-local-images-minikube) for more information.

```zsh
eval $(minikube -p minikube docker-env)  # Point shell to minikube's docker-daemon
minikube image build -t simple-k8s-app:$(date +%Y%m%d%H%M%S) -f ./Dockerfile .  # Build docker image in minikube
minikube image ls --format table  # List images in minikube
# Change image tag in deployment.yaml
kubectl apply -f deployment.yaml  # Apply deployment
kubectl apply -f service.yaml  # Apply service
kubectl get deployments  # Get all deployments
kubectl get services  # Get all services
minikube service simple-k8s-app  # Open service
```

### Apply Changes

```zsh
minikube image build -t simple-k8s-app:$(date +%Y%m%d%H%M%S) -f ./Dockerfile .
minikube image ls --format table
# Update image tag in deployment.yaml
kubectl apply -f deployment.yaml
minikube service simple-k8s-app
```

## Clean Up

```zsh
kubectl delete service simple-k8s-app  # Delete service
kubectl delete deployment simple-k8s-app  # Delete deployment
minikube image rm simple-k8s-app  # Remove image
minikube stop  # Stop minikube
unset DOCKER_HOST DOCKER_CERT_PATH DOCKER_TLS_VERIFY DOCKER_API_VERSION  # Unset docker env
```

## Commands

```zsh
kubectl config get-contexts  # Get all contexts
kubectl config current-context  # Get the current context
kubectl get pods  # Get all pods
minikube stop  # Stop minikube
kubectl logs <pod-name>  # Get logs of a pod
```
