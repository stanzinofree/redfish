# Scorciatoie e Suggerimenti Kubectl

## kubectl aliases
**Tags**: kubectl, alias, scorciatoie, produttività
**Keywords**: alias scorciatoie veloci produttività suggerimenti

Alias comuni per kubectl

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
**Keywords**: completion autocompletamento bash zsh shell setup

Configura autocompletamento shell

```sh
# Bash
source <(kubectl completion bash)
echo "source <(kubectl completion bash)" >> ~/.bashrc

# Zsh
source <(kubectl completion zsh)
echo "source <(kubectl completion zsh)" >> ~/.zshrc

# Completamento alias (bash)
complete -o default -F __start_kubectl k
```

## kubectl explain
**Tags**: kubectl, explain, documentazione, api, riferimento
**Keywords**: explain documentazione api riferimento aiuto campi

Ottieni documentazione per le risorse

```sh
kubectl explain pods
kubectl explain pods.spec
kubectl explain pods.spec.containers
kubectl explain deployment.spec.replicas
```

## kubectl diff
**Tags**: kubectl, diff, confronta, modifiche, anteprima
**Keywords**: diff confronta modifiche anteprima dry-run prova

Anteprima modifiche prima dell'applicazione

```sh
kubectl diff -f deployment.yaml
kubectl diff -k ./cartella-kustomize
```

## kubectl wait
**Tags**: kubectl, wait, condizione, ready, sincronizzazione
**Keywords**: wait aspetta condizione ready sincronizzazione timeout

Attendi una condizione

```sh
kubectl wait --for=condition=ready pod/nome-pod
kubectl wait --for=condition=available deployment/nome-deployment --timeout=60s
kubectl wait --for=delete pod/nome-pod --timeout=30s
```

## kubectl run quick pod
**Tags**: kubectl, run, pod, veloce, test, debug
**Keywords**: run pod veloce test debug temporaneo effimero

Esegui rapidamente un pod di test

```sh
kubectl run test --image=nginx --rm -it -- bash
kubectl run debug --image=busybox --rm -it -- sh
kubectl run curl --image=curlimages/curl --rm -it -- sh
```

## kubectl patch
**Tags**: kubectl, patch, aggiorna, modifica, inline
**Keywords**: patch aggiorna modifica inline veloce cambia

Modifica risorse inline

```sh
kubectl patch deployment nome-deployment -p '{"spec":{"replicas":3}}'
kubectl patch pod nome-pod -p '{"metadata":{"labels":{"env":"prod"}}}'
```

## kubectl label
**Tags**: kubectl, label, metadata, organizza, seleziona
**Keywords**: label etichetta metadata organizza seleziona tag

Aggiungi o modifica etichette

```sh
kubectl label pods nome-pod env=prod
kubectl label pods nome-pod env-
kubectl label pods --all env=staging
```

## kubectl annotate
**Tags**: kubectl, annotate, metadata, documentazione
**Keywords**: annotate annotazione metadata documentazione note

Aggiungi o modifica annotazioni

```sh
kubectl annotate pods nome-pod description="pod di test"
kubectl annotate pods nome-pod description-
```

## kubectl drain node
**Tags**: kubectl, drain, nodo, manutenzione, evict
**Keywords**: drain svuota nodo manutenzione evict cordon uncordon

Svuota nodo per manutenzione

```sh
kubectl drain nome-nodo --ignore-daemonsets
kubectl cordon nome-nodo
kubectl uncordon nome-nodo
```

## kubectl debug
**Tags**: kubectl, debug, troubleshoot, effimero, contenitore
**Keywords**: debug troubleshoot diagnostica effimero contenitore diagnosi

Debug pod in esecuzione

```sh
kubectl debug nome-pod -it --image=busybox
kubectl debug nome-pod -it --image=ubuntu --target=nome-contenitore
kubectl debug node/nome-nodo -it --image=ubuntu
```

## kubectl events
**Tags**: kubernetes, kubectl, eventi, logs, cluster, troubleshoot
**Keywords**: eventi logs cluster troubleshoot cronologia diagnostica

Visualizza eventi del cluster

```sh
kubectl get events
kubectl get events --sort-by=.metadata.creationTimestamp
kubectl get events -n nome-namespace
kubectl get events --field-selector type=Warning
```
