# multipass-skaffold-demo

Pre-reqs for the demo:
* git
* multipass
  
Demo Steps:
1. Clone this repo and change directory to the repo folder
2. Run multipass-setup.sh
3. Start up minikube
4. Launch skaffold via the command:   skaffold dev
5. After build and deploy complete
   - Check k8s deployment:   kubectl get deploy,po
   - Load browser:   http://localhost:80/test
6. Make changes to any of the code - Dockerfile or main.go
7. Skaffold will notice the changes, rebuild & deploy
   - Check k8s deployment again:   kubectl get deploy,po 
   - Load browser again:  http://localhost:80/test
8. Ctrl+c to exit skaffold (this will tear down the k8s artifacts)
