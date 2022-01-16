#!/bin/sh

# let's start minikube with 4CPUs. i'll use virtualbox driver here(not needed for ubuntu)
# and the host-onky-cidr is CIDR or network which your virtual machine running docker will connect to.
# ref:https://www.virtualbox.org/manual/ch06.html#network_hostonly
minikube start --driver=docker --cpus=4 --host-only-cidr=192.168.77.1/24
##
### let's say to minikube to re-use the Docker daemon inside the Minikube instance:
eval $(minikube docker-env)
#
## building docker image (you can skip it for sure)
docker build --rm -t be-router .
#
## Create mongodb service with mongodb stateful-set
kubectl apply -f k8s/be-router/configMaps.yaml
kubectl apply -f k8s/be-router/be-router-volumes.yaml
sleep 15
kubectl apply -f k8s/be-router/be-router-redis.yaml
kubectl apply -f k8s/be-router/be-router-mongo.yaml
kubectl apply -f k8s/be-router/be-router-mysql.yaml
kubectl apply -f k8s/be-router/be-router.yaml
sleep 5
#
## Print current deployment state (unlikely to be finished yet)
kubectl get all

## Print current deployment state (unlikely to be finished yet)
kubectl get all

echo "As you can see, external IP for be-router app is pending, please open a new terminal and run there minikube tunnel \nPress Ctrl-C in the terminal can be used to terminate the process at which time the network routes will be cleaned up "
##
# Using minikube tunnel
# Services of type LoadBalancer can be exposed via the minikube tunnel command.
# It must be run in a separate terminal window to keep the LoadBalancer running.
# Ctrl-C in the terminal can be used to terminate the process at which time the network routes will be cleaned up

##not need to run it everytime, but it won't hurt
chmod +x stop.sh
