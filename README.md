# multipass-skaffold-demo

Pre-reqs for the demo:
* git
* multipass
  
Demo Steps:
1.  Clone this repo and change directory to the repo folder
```
git clone https://github.com/MxNxPx/multipass-skaffold-demo.git && cd $_
```
2.  Spin up a multipass instance with prereqs
```
bash multipass-setup.sh
```
3.  Shell into the multipass instance
```
multipass shell ubuntu-multipass
```
4.  Start up minikube and enable ingress
```
sudo minikube start --vm-driver=none --kubernetes-version v1.16.3
sudo minikube addons enable ingress
sudo chown -R $USER:$USER ~/.minikube/ ~/.kube/
```
5.  Launch skaffold
```
cd ~/skaffold
sudo skaffold dev
```
6.  Open another shell into multipass and get the service URL
```
mutlipass shell ubuntu-multipass
sudo minikube service cowsay-server --url
```
7.  Validate working
    - Check k8s deployment:  `sudo kubectl get deploy,po`
    - Load browser:  (URL from above)

8.  Make changes to any of the code - Dockerfile or main.go

9.  Skaffold will notice the changes, rebuild & deploy
    - Check k8s deployment again:  `sudo kubectl get deploy,po` 
    - Load browser again:  (URL from above)

10. Ctrl+c to exit skaffold (this will tear down the k8s artifacts)

11. Completely tear down the demo (when you are done)
```
multipass delete ubuntu-multipass
multipass purge
```
