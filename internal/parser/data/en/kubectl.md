# Kubectl Shortcuts and Tips

## kubectl aliases
**Tags**: kubectl, alias, shortcuts, productivity
**Keywords**: alias shortcuts quick productivity tips

Common kubectl aliases

```sh
alias k=kubectl
alias kgp='kubectl get pods'
alias kgs='kubectl get services'
alias kgd='kubectl get deployments'
alias kga='kubectl get all'
alias kaf='kubectl apply -f'
alias kdel='kubectl delete'
alias kl='kubectl logs'
alias kex='kubectl exec -it'
```

## kubectl completion
**Tags**: kubectl, completion, bash, zsh, shell
**Keywords**: completion autocomplete bash zsh shell setup

Set up shell completion

```sh
# Bash
source <(kubectl completion bash)
echo "source <(kubectl completion bash)" >> ~/.bashrc

# Zsh
source <(kubectl completion zsh)
echo "source <(kubectl completion zsh)" >> ~/.zshrc

# Alias completion (bash)
complete -o default -F __start_kubectl k
```

## kubectl explain
**Tags**: kubectl, explain, documentation, api, reference
**Keywords**: explain documentation api reference help fields

Get documentation for resources

```sh
kubectl explain pods
kubectl explain pods.spec
kubectl explain pods.spec.containers
kubectl explain deployment.spec.replicas
```

## kubectl diff
**Tags**: kubectl, diff, compare, changes, preview
**Keywords**: diff compare changes preview dry-run

Preview changes before applying

```sh
kubectl diff -f deployment.yaml
kubectl diff -k ./kustomize-dir
```

## kubectl wait
**Tags**: kubectl, wait, condition, ready, synchronization
**Keywords**: wait condition ready synchronization timeout

Wait for a condition

```sh
kubectl wait --for=condition=ready pod/pod-name
kubectl wait --for=condition=available deployment/deployment-name --timeout=60s
kubectl wait --for=delete pod/pod-name --timeout=30s
```

## kubectl run quick pod
**Tags**: kubectl, run, pod, quick, test, debug
**Keywords**: run pod quick test debug temporary ephemeral

Quickly run a test pod

```sh
kubectl run test --image=nginx --rm -it -- bash
kubectl run debug --image=busybox --rm -it -- sh
kubectl run curl --image=curlimages/curl --rm -it -- sh
```

## kubectl patch
**Tags**: kubectl, patch, update, modify, inline
**Keywords**: patch update modify inline quick change

Patch resources inline

```sh
kubectl patch deployment deployment-name -p '{"spec":{"replicas":3}}'
kubectl patch pod pod-name -p '{"metadata":{"labels":{"env":"prod"}}}'
```

## kubectl label
**Tags**: kubectl, label, metadata, organize, select
**Keywords**: label metadata organize select tag

Add or modify labels

```sh
kubectl label pods pod-name env=prod
kubectl label pods pod-name env-
kubectl label pods --all env=staging
```

## kubectl annotate
**Tags**: kubectl, annotate, metadata, documentation
**Keywords**: annotate metadata documentation notes

Add or modify annotations

```sh
kubectl annotate pods pod-name description="test pod"
kubectl annotate pods pod-name description-
```

## kubectl drain node
**Tags**: kubectl, drain, node, maintenance, evict
**Keywords**: drain node maintenance evict cordon uncordon

Drain node for maintenance

```sh
kubectl drain node-name --ignore-daemonsets
kubectl cordon node-name
kubectl uncordon node-name
```

## kubectl debug
**Tags**: kubectl, debug, troubleshoot, ephemeral, container
**Keywords**: debug troubleshoot ephemeral container diagnose

Debug running pods

```sh
kubectl debug pod-name -it --image=busybox
kubectl debug pod-name -it --image=ubuntu --target=container-name
kubectl debug node/node-name -it --image=ubuntu
```

## kubectl events
**Tags**: kubectl, events, logs, cluster, troubleshoot
**Keywords**: events logs cluster troubleshoot history

View cluster events

```sh
kubectl get events
kubectl get events --sort-by=.metadata.creationTimestamp
kubectl get events -n namespace-name
kubectl get events --field-selector type=Warning
```
