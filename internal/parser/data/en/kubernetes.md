# Kubernetes Commands

## kubectl get pods
**Tags**: kubernetes, kubectl, pods, list, status
**Keywords**: get pods list containers running status show
**Short_Description**: List all pods in the current namespace
**Long_Description**: Displays information about pods including their name, ready status, number of restarts, age, and status. Use -A to list pods from all namespaces, -n to specify a namespace, and -o wide for additional details like IP and node. Essential for monitoring pod health.

```sh
kubectl get pods
kubectl get pods -A
kubectl get pods -n namespace-name
kubectl get pods -o wide
```

## kubectl describe pod
**Tags**: kubernetes, kubectl, pod, describe, details, debug
**Keywords**: describe pod details information debug troubleshoot
**Short_Description**: Show detailed information about a pod
**Long_Description**: Provides comprehensive information about a specific pod including events, container status, resource requests/limits, volumes, and conditions. Critical for debugging pod issues, understanding failures, and viewing recent events that affected the pod.

```sh
kubectl describe pod pod-name
kubectl describe pod pod-name -n namespace-name
```

## kubectl logs
**Tags**: kubernetes, kubectl, logs, debugging, output
**Keywords**: logs debug output tail follow container
**Short_Description**: View pod logs
**Long_Description**: Retrieves logs from containers running in a pod. The -f flag follows logs in real-time, -c specifies a container in multi-container pods, and --tail limits output to last N lines. Essential for application debugging and monitoring.

```sh
kubectl logs pod-name
kubectl logs -f pod-name
kubectl logs pod-name -c container-name
kubectl logs --tail=100 pod-name
```

## kubectl exec
**Tags**: kubernetes, kubectl, exec, shell, command, debug
**Keywords**: exec shell bash command execute debug interactive
**Short_Description**: Execute command in a pod
**Long_Description**: Runs commands inside a running pod container. The -it flags provide interactive terminal access, commonly used for opening a shell. Use -c to specify container in multi-container pods. The -- separates kubectl flags from the command to execute.

```sh
kubectl exec -it pod-name -- bash
kubectl exec -it pod-name -- sh
kubectl exec pod-name -- ls /app
kubectl exec -it pod-name -c container-name -- bash
```

## kubectl apply
**Tags**: kubernetes, kubectl, apply, deploy, create, manifest
**Keywords**: apply deploy create manifest yaml configuration
**Short_Description**: Apply configuration from file
**Long_Description**: Creates or updates Kubernetes resources from YAML/JSON manifests. Supports files, URLs, and Kustomize directories with -k flag. Uses declarative approach - safe to run multiple times as it only applies changes. Preferred method for GitOps workflows.

```sh
kubectl apply -f deployment.yaml
kubectl apply -f https://example.com/manifest.yaml
kubectl apply -k ./kustomize-dir
```

## kubectl delete
**Tags**: kubernetes, kubectl, delete, remove, cleanup
**Keywords**: delete remove cleanup destroy resource
**Short_Description**: Delete resources
**Long_Description**: Removes Kubernetes resources by name, file, or label selector. Can delete specific resources or all resources of a type using --all. Supports graceful termination with configurable grace period. Use with caution in production environments.

```sh
kubectl delete pod pod-name
kubectl delete -f deployment.yaml
kubectl delete deployment deployment-name
kubectl delete pods --all
```

## kubectl get services
**Tags**: kubernetes, kubectl, services, svc, networking, list
**Keywords**: get services svc list networking endpoints
**Short_Description**: List services
**Long_Description**: Shows Kubernetes services that expose applications running in pods. Displays service name, type (ClusterIP, NodePort, LoadBalancer), cluster IP, external IP, ports, and age. Use -A for all namespaces or -o wide for additional details.

```sh
kubectl get services
kubectl get svc
kubectl get svc -A
kubectl get svc -o wide
```

## kubectl get deployments
**Tags**: kubernetes, kubectl, deployments, list, apps
**Keywords**: get deployments list applications running status
**Short_Description**: List deployments
**Long_Description**: Displays deployments showing desired vs ready replicas, up-to-date status, available replicas, and age. Deployments manage ReplicaSets which in turn manage pods. Essential for monitoring application deployment status and health.

```sh
kubectl get deployments
kubectl get deploy
kubectl get deployments -A
kubectl get deployments -o wide
```

## kubectl scale
**Tags**: kubernetes, kubectl, scale, replicas, horizontal
**Keywords**: scale replicas horizontal autoscale increase decrease
**Short_Description**: Scale deployment replicas
**Long_Description**: Adjusts the number of pod replicas in a deployment, replicaset, or statefulset. Increases or decreases replicas to handle load changes. Can scale from file or by resource name. Changes take effect immediately but pod creation/termination follows Kubernetes scheduling.

```sh
kubectl scale deployment deployment-name --replicas=3
kubectl scale --replicas=5 -f deployment.yaml
```

## kubectl rollout
**Tags**: kubernetes, kubectl, rollout, deployment, update, rollback
**Keywords**: rollout deployment update rollback restart history
**Short_Description**: Manage deployment rollouts
**Long_Description**: Controls deployment updates including status monitoring, viewing revision history, rolling back to previous versions, and restarting deployments. Essential for managing application updates safely and recovering from failed deployments.

```sh
kubectl rollout status deployment/deployment-name
kubectl rollout history deployment/deployment-name
kubectl rollout undo deployment/deployment-name
kubectl rollout restart deployment/deployment-name
```

## kubectl port-forward
**Tags**: kubernetes, kubectl, port-forward, networking, debug
**Keywords**: port forward tunnel local debugging access
**Short_Description**: Forward local port to pod
**Long_Description**: Creates a tunnel from local machine to a pod, service, or deployment in the cluster. Maps local port to container port for direct access without exposing services publicly. Useful for debugging, testing, and accessing cluster resources locally.

```sh
kubectl port-forward pod-name 8080:80
kubectl port-forward svc/service-name 8080:80
kubectl port-forward deployment/deployment-name 8080:80
```

## kubectl create namespace
**Tags**: kubernetes, kubectl, namespace, create, organization
**Keywords**: create namespace new organization isolation
**Short_Description**: Create a new namespace
**Long_Description**: Creates a new namespace for logical isolation of resources within the cluster. Namespaces provide scope for resource names, enable resource quotas and access control, and help organize multi-tenant or multi-environment clusters.

```sh
kubectl create namespace namespace-name
kubectl create ns namespace-name
```

## kubectl get nodes
**Tags**: kubernetes, kubectl, nodes, cluster, infrastructure, list
**Keywords**: get nodes list cluster infrastructure workers masters
**Short_Description**: List cluster nodes
**Long_Description**: Displays cluster nodes showing their status (Ready/NotReady), roles (control-plane, worker), age, and Kubernetes version. Use -o wide for additional details like internal/external IPs, OS, kernel, and container runtime. Critical for cluster health monitoring.

```sh
kubectl get nodes
kubectl get nodes -o wide
kubectl describe node node-name
```

## kubectl top
**Tags**: kubernetes, kubectl, metrics, resources, cpu, memory
**Keywords**: top metrics resources cpu memory usage statistics
**Short_Description**: Show resource usage
**Long_Description**: Displays current CPU and memory usage for nodes or pods. Requires metrics-server to be installed in the cluster. Use --containers flag to show individual container metrics within pods. Essential for monitoring resource consumption and identifying performance issues.

```sh
kubectl top nodes
kubectl top pods
kubectl top pod pod-name --containers
```

## kubectl config
**Tags**: kubernetes, kubectl, config, context, cluster, kubeconfig
**Keywords**: config context cluster kubeconfig switch current
**Short_Description**: Manage kubeconfig
**Long_Description**: Manages kubectl configuration including contexts, clusters, and users. View current config, list available contexts, switch between clusters/namespaces, and set default namespace. Contexts combine cluster, user, and namespace information for easy switching between environments.

```sh
kubectl config view
kubectl config get-contexts
kubectl config use-context context-name
kubectl config current-context
kubectl config set-context --current --namespace=namespace-name
```
