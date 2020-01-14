# cowsay-server demo

Pre-reqs for the demo:
* kubectl - https://kubernetes.io/docs/tasks/tools/install-kubectl/
* skaffold - https://github.com/GoogleContainerTools/skaffold/releases
* k8s cluster
  - Docker for Desktop
  - Minikube
  
Demo Steps:
1. Clone this repo and change directory to the cowsay-server folder
2. Confirm you are pointed at the desired k8s cluster:   kubectl config current-context
3. Launch skaffold via the command:   skaffold dev
4. After build and deploy complete
   - Check k8s deployment:   kubectl get deploy,po
   - Load browser:   http://localhost:80/test
5. Make changes to any of the code - Dockerfile or main.go
6. Skaffold will notice the changes, rebuild & deploy
   - Check k8s deployment again:   kubectl get deploy,po 
   - Load browser again:  http://localhost:80/test
7. Ctrl+c to exit skaffold (this will tear down the k8s artifacts)
