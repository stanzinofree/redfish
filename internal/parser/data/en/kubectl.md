# Kubectl Shortcuts and Tips

## kubectl aliases
**Tags**: kubectl, alias, shortcuts, productivity
**Keywords**: alias shortcuts quick productivity tips
**Short_Description**: Common kubectl aliases
**Long_Description**: Set up shell aliases to reduce typing and speed up kubectl workflow. These aliases create shortcuts for frequently used kubectl commands. Add them to your .bashrc or .zshrc for permanent availability. Improves productivity significantly for daily Kubernetes operations.

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
**Short_Description**: Set up shell completion
**Long_Description**: Enables tab completion for kubectl commands, flags, and resource names in your shell. Supports bash, zsh, and other shells. Significantly speeds up command entry and reduces errors by auto-completing resource names and flags. Must be configured once per shell.

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
**Short_Description**: Get documentation for resources
**Long_Description**: Displays documentation for Kubernetes resource fields directly from the API schema. Navigate nested fields with dot notation. Essential for understanding available fields, their types, and requirements when writing manifests. Works offline using local schema.

```sh
kubectl explain pods
kubectl explain pods.spec
kubectl explain pods.spec.containers
kubectl explain deployment.spec.replicas
```

## kubectl diff
**Tags**: kubectl, diff, compare, changes, preview
**Keywords**: diff compare changes preview dry-run
**Short_Description**: Preview changes before applying
**Long_Description**: Shows what changes would be made by kubectl apply before actually applying them. Compares local manifests with live cluster state and displays differences. Essential for safe deployments - review changes before they affect production. Works with files and Kustomize directories.

```sh
kubectl diff -f deployment.yaml
kubectl diff -k ./kustomize-dir
```

## kubectl wait
**Tags**: kubectl, wait, condition, ready, synchronization
**Keywords**: wait condition ready synchronization timeout
**Short_Description**: Wait for a condition
**Long_Description**: Blocks until a specified condition is met on Kubernetes resources. Useful in scripts and CI/CD pipelines to ensure resources are ready before proceeding. Supports conditions like ready, available, deleted with configurable timeouts. Returns exit code on timeout.

```sh
kubectl wait --for=condition=ready pod/pod-name
kubectl wait --for=condition=available deployment/deployment-name --timeout=60s
kubectl wait --for=delete pod/pod-name --timeout=30s
```

## kubectl run quick pod
**Tags**: kubectl, run, pod, quick, test, debug
**Keywords**: run pod quick test debug temporary ephemeral
**Short_Description**: Quickly run a test pod
**Long_Description**: Creates an ephemeral pod for quick testing or debugging. The --rm flag deletes the pod on exit, -it provides interactive terminal. Useful for testing network connectivity, running quick commands, or debugging cluster issues without creating permanent resources.

```sh
kubectl run test --image=nginx --rm -it -- bash
kubectl run debug --image=busybox --rm -it -- sh
kubectl run curl --image=curlimages/curl --rm -it -- sh
```

## kubectl patch
**Tags**: kubectl, patch, update, modify, inline
**Keywords**: patch update modify inline quick change
**Short_Description**: Patch resources inline
**Long_Description**: Updates specific fields of resources using JSON/YAML patch without needing full manifest files. Useful for quick changes to running resources. Supports strategic merge patch, JSON merge patch, and JSON patch. Faster than edit for scripted updates.

```sh
kubectl patch deployment deployment-name -p '{"spec":{"replicas":3}}'
kubectl patch pod pod-name -p '{"metadata":{"labels":{"env":"prod"}}}'
```

## kubectl label
**Tags**: kubectl, label, metadata, organize, select
**Keywords**: label metadata organize select tag
**Short_Description**: Add or modify labels
**Long_Description**: Manages labels on Kubernetes resources. Labels are key-value pairs used for organizing, selecting, and filtering resources. Add labels with key=value, remove with key-. Can apply to multiple resources using --all or selectors. Critical for service discovery and deployments.

```sh
kubectl label pods pod-name env=prod
kubectl label pods pod-name env-
kubectl label pods --all env=staging
```

## kubectl annotate
**Tags**: kubectl, annotate, metadata, documentation
**Keywords**: annotate metadata documentation notes
**Short_Description**: Add or modify annotations
**Long_Description**: Manages annotations on resources. Annotations store non-identifying metadata like build info, deployment notes, or tool-specific data. Unlike labels, they're not used for selection. Remove annotations with key-. Useful for documentation and integration with external tools.

```sh
kubectl annotate pods pod-name description="test pod"
kubectl annotate pods pod-name description-
```

## kubectl drain node
**Tags**: kubectl, drain, node, maintenance, evict
**Keywords**: drain node maintenance evict cordon uncordon
**Short_Description**: Drain node for maintenance
**Long_Description**: Safely evicts pods from a node for maintenance. Cordon prevents new pods from scheduling, drain evicts existing pods respecting PodDisruptionBudgets. Use --ignore-daemonsets to ignore DaemonSet pods. Uncordon marks node schedulable again after maintenance.

```sh
kubectl drain node-name --ignore-daemonsets
kubectl cordon node-name
kubectl uncordon node-name
```

## kubectl debug
**Tags**: kubectl, debug, troubleshoot, ephemeral, container
**Keywords**: debug troubleshoot ephemeral container diagnose
**Short_Description**: Debug running pods
**Long_Description**: Creates ephemeral debug containers in running pods or nodes without modifying the pod specification. Useful for debugging distroless or minimal images that lack debugging tools. Can attach to existing containers or create new debug containers with full toolsets.

```sh
kubectl debug pod-name -it --image=busybox
kubectl debug pod-name -it --image=ubuntu --target=container-name
kubectl debug node/node-name -it --image=ubuntu
```

## kubectl events
**Tags**: kubectl, events, logs, cluster, troubleshoot
**Keywords**: events logs cluster troubleshoot history
**Short_Description**: View cluster events
**Long_Description**: Displays cluster-wide events showing state changes, errors, and warnings. Events are time-limited (usually 1 hour) and provide insights into cluster operations. Sort by timestamp, filter by namespace or type. Essential for troubleshooting pod failures and cluster issues.

```sh
kubectl get events
kubectl get events --sort-by=.metadata.creationTimestamp
kubectl get events -n namespace-name
kubectl get events --field-selector type=Warning
```
