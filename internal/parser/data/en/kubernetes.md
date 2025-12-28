# Kubernetes Commands

## kubectl get pods
**Tags**: kubernetes, kubectl, pods, list, status
**Keywords**: get pods list containers running status show

List all pods in the current namespace

```sh
kubectl get pods
kubectl get pods -A
kubectl get pods -n namespace-name
kubectl get pods -o wide
```

## kubectl describe pod
**Tags**: kubernetes, kubectl, pod, describe, details, debug
**Keywords**: describe pod details information debug troubleshoot

Show detailed information about a pod

```sh
kubectl describe pod pod-name
kubectl describe pod pod-name -n namespace-name
```

## kubectl logs
**Tags**: kubernetes, kubectl, logs, debugging, output
**Keywords**: logs debug output tail follow container

View pod logs

```sh
kubectl logs pod-name
kubectl logs -f pod-name
kubectl logs pod-name -c container-name
kubectl logs --tail=100 pod-name
```

## kubectl exec
**Tags**: kubernetes, kubectl, exec, shell, command, debug
**Keywords**: exec shell bash command execute debug interactive

Execute command in a pod

```sh
kubectl exec -it pod-name -- bash
kubectl exec -it pod-name -- sh
kubectl exec pod-name -- ls /app
kubectl exec -it pod-name -c container-name -- bash
```

## kubectl apply
**Tags**: kubernetes, kubectl, apply, deploy, create, manifest
**Keywords**: apply deploy create manifest yaml configuration

Apply configuration from file

```sh
kubectl apply -f deployment.yaml
kubectl apply -f https://example.com/manifest.yaml
kubectl apply -k ./kustomize-dir
```

## kubectl delete
**Tags**: kubernetes, kubectl, delete, remove, cleanup
**Keywords**: delete remove cleanup destroy resource

Delete resources

```sh
kubectl delete pod pod-name
kubectl delete -f deployment.yaml
kubectl delete deployment deployment-name
kubectl delete pods --all
```

## kubectl get services
**Tags**: kubernetes, kubectl, services, svc, networking, list
**Keywords**: get services svc list networking endpoints

List services

```sh
kubectl get services
kubectl get svc
kubectl get svc -A
kubectl get svc -o wide
```

## kubectl get deployments
**Tags**: kubernetes, kubectl, deployments, list, apps
**Keywords**: get deployments list applications running status

List deployments

```sh
kubectl get deployments
kubectl get deploy
kubectl get deployments -A
kubectl get deployments -o wide
```

## kubectl scale
**Tags**: kubernetes, kubectl, scale, replicas, horizontal
**Keywords**: scale replicas horizontal autoscale increase decrease

Scale deployment replicas

```sh
kubectl scale deployment deployment-name --replicas=3
kubectl scale --replicas=5 -f deployment.yaml
```

## kubectl rollout
**Tags**: kubernetes, kubectl, rollout, deployment, update, rollback
**Keywords**: rollout deployment update rollback restart history

Manage deployment rollouts

```sh
kubectl rollout status deployment/deployment-name
kubectl rollout history deployment/deployment-name
kubectl rollout undo deployment/deployment-name
kubectl rollout restart deployment/deployment-name
```

## kubectl port-forward
**Tags**: kubernetes, kubectl, port-forward, networking, debug
**Keywords**: port forward tunnel local debugging access

Forward local port to pod

```sh
kubectl port-forward pod-name 8080:80
kubectl port-forward svc/service-name 8080:80
kubectl port-forward deployment/deployment-name 8080:80
```

## kubectl create namespace
**Tags**: kubernetes, kubectl, namespace, create, organization
**Keywords**: create namespace new organization isolation

Create a new namespace

```sh
kubectl create namespace namespace-name
kubectl create ns namespace-name
```

## kubectl get nodes
**Tags**: kubernetes, kubectl, nodes, cluster, infrastructure, list
**Keywords**: get nodes list cluster infrastructure workers masters

List cluster nodes

```sh
kubectl get nodes
kubectl get nodes -o wide
kubectl describe node node-name
```

## kubectl top
**Tags**: kubernetes, kubectl, metrics, resources, cpu, memory
**Keywords**: top metrics resources cpu memory usage statistics

Show resource usage

```sh
kubectl top nodes
kubectl top pods
kubectl top pod pod-name --containers
```

## kubectl config
**Tags**: kubernetes, kubectl, config, context, cluster, kubeconfig
**Keywords**: config context cluster kubeconfig switch current

Manage kubeconfig

```sh
kubectl config view
kubectl config get-contexts
kubectl config use-context context-name
kubectl config current-context
kubectl config set-context --current --namespace=namespace-name
```
