# Comandi Kubernetes

## kubectl get pods
**Tags**: kubernetes, kubectl, pods, lista, stato
**Keywords**: get pods lista contenitori running stato mostra visualizza

Lista tutti i pod nel namespace corrente

```sh
kubectl get pods
kubectl get pods -A
kubectl get pods -n nome-namespace
kubectl get pods -o wide
```

## kubectl describe pod
**Tags**: kubernetes, kubectl, pod, describe, dettagli, debug
**Keywords**: describe pod dettagli informazioni debug troubleshoot diagnostica

Mostra informazioni dettagliate su un pod

```sh
kubectl describe pod nome-pod
kubectl describe pod nome-pod -n nome-namespace
```

## kubectl logs
**Tags**: kubernetes, kubectl, logs, debugging, output
**Keywords**: logs debug output tail follow contenitore

Visualizza i log dei pod

```sh
kubectl logs nome-pod
kubectl logs -f nome-pod
kubectl logs nome-pod -c nome-contenitore
kubectl logs --tail=100 nome-pod
```

## kubectl exec
**Tags**: kubernetes, kubectl, exec, shell, comando, debug
**Keywords**: exec shell bash comando esegui debug interattivo

Esegui comando in un pod

```sh
kubectl exec -it nome-pod -- bash
kubectl exec -it nome-pod -- sh
kubectl exec nome-pod -- ls /app
kubectl exec -it nome-pod -c nome-contenitore -- bash
```

## kubectl apply
**Tags**: kubernetes, kubectl, apply, deploy, crea, manifest
**Keywords**: apply deploy crea manifest yaml configurazione

Applica configurazione da file

```sh
kubectl apply -f deployment.yaml
kubectl apply -f https://example.com/manifest.yaml
kubectl apply -k ./cartella-kustomize
```

## kubectl delete
**Tags**: kubernetes, kubectl, delete, rimuovi, cleanup
**Keywords**: delete rimuovi elimina cleanup distruggi risorsa

Elimina risorse

```sh
kubectl delete pod nome-pod
kubectl delete -f deployment.yaml
kubectl delete deployment nome-deployment
kubectl delete pods --all
```

## kubectl get services
**Tags**: kubernetes, kubectl, services, svc, networking, lista
**Keywords**: get services svc lista networking endpoint

Lista servizi

```sh
kubectl get services
kubectl get svc
kubectl get svc -A
kubectl get svc -o wide
```

## kubectl get deployments
**Tags**: kubernetes, kubectl, deployments, lista, apps
**Keywords**: get deployments lista applicazioni running stato

Lista deployment

```sh
kubectl get deployments
kubectl get deploy
kubectl get deployments -A
kubectl get deployments -o wide
```

## kubectl scale
**Tags**: kubernetes, kubectl, scale, repliche, orizzontale
**Keywords**: scale repliche orizzontale autoscale aumenta diminuisci

Scala le repliche del deployment

```sh
kubectl scale deployment nome-deployment --replicas=3
kubectl scale --replicas=5 -f deployment.yaml
```

## kubectl rollout
**Tags**: kubernetes, kubectl, rollout, deployment, aggiorna, rollback
**Keywords**: rollout deployment aggiorna rollback riavvia cronologia

Gestisci i rollout dei deployment

```sh
kubectl rollout status deployment/nome-deployment
kubectl rollout history deployment/nome-deployment
kubectl rollout undo deployment/nome-deployment
kubectl rollout restart deployment/nome-deployment
```

## kubectl port-forward
**Tags**: kubernetes, kubectl, port-forward, networking, debug
**Keywords**: port forward tunnel locale debugging accesso

Inoltra porta locale al pod

```sh
kubectl port-forward nome-pod 8080:80
kubectl port-forward svc/nome-servizio 8080:80
kubectl port-forward deployment/nome-deployment 8080:80
```

## kubectl create namespace
**Tags**: kubernetes, kubectl, namespace, crea, organizzazione
**Keywords**: crea namespace nuovo organizzazione isolamento

Crea un nuovo namespace

```sh
kubectl create namespace nome-namespace
kubectl create ns nome-namespace
```

## kubectl get nodes
**Tags**: kubernetes, kubectl, nodes, cluster, infrastruttura, lista
**Keywords**: get nodes lista cluster infrastruttura worker master

Lista nodi del cluster

```sh
kubectl get nodes
kubectl get nodes -o wide
kubectl describe node nome-nodo
```

## kubectl top
**Tags**: kubernetes, kubectl, metriche, risorse, cpu, memoria
**Keywords**: top metriche risorse cpu memoria utilizzo statistiche

Mostra utilizzo risorse

```sh
kubectl top nodes
kubectl top pods
kubectl top pod nome-pod --containers
```

## kubectl config
**Tags**: kubernetes, kubectl, config, contesto, cluster, kubeconfig
**Keywords**: config contesto cluster kubeconfig switch corrente cambia

Gestisci kubeconfig

```sh
kubectl config view
kubectl config get-contexts
kubectl config use-context nome-contesto
kubectl config current-context
kubectl config set-context --current --namespace=nome-namespace
```
