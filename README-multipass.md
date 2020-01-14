sudo minikube start --vm-driver=none --kubernetes-version v1.16.3
sudo minikube addons enable ingress
sudo chown -R $USER:$USER ~/.minikube/ ~/.kube/

cd ~/multipass
skaffold dev

minikube service cowsay-server --url
