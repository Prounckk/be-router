#!/bin/sh
# Stop all
kubectl delete -f k8s/be-router/configMaps.yaml
kubectl delete -f k8s/be-router/be-router-volumes.yaml
kubectl delete -f k8s/be-router/be-router-redis.yaml
kubectl delete -f k8s/be-router/be-router-mongo.yaml
kubectl delete -f k8s/be-router/be-router.yaml

# Stop minikube
minikube stop
