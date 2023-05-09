Practing running services in kube. 

## What is it?

A frontend of Javascript (react?) served up with nginx
A middletier web API that routes requests to and from the backend (Java Spring)
A backend API that takes input and returns an output. (Python)

Plus all the docker and kube manifests required to build it and run it in kubernetes.


## KinD setup

```
cd kind
make create

```


I hardcoded some URLs to make it easy to run locally while simulating a real world scnenario.

Add the following to your /etc/hosts file:
   127.0.0.1 sa.info

   127.0.0.1 api.sa.info



## Launch the app

    kubectl apply -f resource-manifests/

Visit the app at http://sa.info
