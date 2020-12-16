A small microservices app I slapped together to help learn kube networking. 

## What is it?

A frontend of Javascript (react?) served up with nginx
A middletier web API that routes requests to and from the backend (Java Spring)
A backend API that takes input and returns an output. (Python)

Plus all the docker and kube manifests required to build it and run it in kubernetes.


## minikube setup

Install minikube with brew:
    brew install minikube

Start minikube
    minikube start

Enable ingress controllers for minikube:
    minikube addons enable ingress

My minikube version was 1.15.1
Kubernetes version 1.18


I hardcoded some URLs to make it easy to run locally while simulating a real world scnenario.

Run
    minikube ip

to get the IP address of your local minikube. The add the following to your /etc/hosts file:
    <minikube address> sa.info
    <minikube address> api.sa.info

If you ever reinstall minikube you might need to double check this IP address to make sure it hasnt changed.

The app uses the Kube LoadBalancer service. To make these work with Minikube you must run 

    minikube tunnel

in a seperate terminal. (it will ask for your login password)



## Launch the app

    kubectl apply -f resource-manifests/

Visit the app at http://sa.info
