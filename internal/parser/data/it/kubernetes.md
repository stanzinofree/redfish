# Comandi Kubernetes

## kubectl get pods
**Tags**: kubernetes, kubectl, pods, lista, stato
**Keywords**: get pods lista contenitori running stato mostra visualizza
**Short_Description**: Lista tutti i pod nel namespace corrente
**Long_Description**: Mostra informazioni sui pod inclusi nome, stato di pronto, numero di riavvii, età e stato. Usa -A per listare pod da tutti i namespace, -n per specificare un namespace, e -o wide per dettagli aggiuntivi come IP e nodo. Essenziale per monitorare la salute dei pod.

```sh
kubectl get pods
kubectl get pods -A
kubectl get pods -n nome-namespace
kubectl get pods -o wide
```

## kubectl describe pod
**Tags**: kubernetes, kubectl, pod, describe, dettagli, debug
**Keywords**: describe pod dettagli informazioni debug troubleshoot diagnostica
**Short_Description**: Mostra informazioni dettagliate su un pod
**Long_Description**: Fornisce informazioni complete su un pod specifico inclusi eventi, stato container, richieste/limiti risorse, volumi e condizioni. Cruciale per il debug di problemi dei pod, comprendere i fallimenti e visualizzare eventi recenti che hanno influenzato il pod.

```sh
kubectl describe pod nome-pod
kubectl describe pod nome-pod -n nome-namespace
```

## kubectl logs
**Tags**: kubernetes, kubectl, logs, debugging, output
**Keywords**: logs debug output tail follow contenitore registro
**Short_Description**: Visualizza i log dei pod
**Long_Description**: Recupera i log dai container in esecuzione in un pod. Il flag -f segue i log in tempo reale, -c specifica un container in pod multi-container, e --tail limita l'output alle ultime N righe. Essenziale per il debug e il monitoraggio delle applicazioni.

```sh
kubectl logs nome-pod
kubectl logs -f nome-pod
kubectl logs nome-pod -c nome-contenitore
kubectl logs --tail=100 nome-pod
```

## kubectl exec
**Tags**: kubernetes, kubectl, exec, shell, comando, debug
**Keywords**: exec shell bash comando esegui debug interattivo eseguire
**Short_Description**: Esegui comando in un pod
**Long_Description**: Esegue comandi dentro un container in esecuzione di un pod. I flag -it forniscono accesso al terminale interattivo, comunemente usato per aprire una shell. Usa -c per specificare il container in pod multi-container. Il -- separa i flag kubectl dal comando da eseguire.

```sh
kubectl exec -it nome-pod -- bash
kubectl exec -it nome-pod -- sh
kubectl exec nome-pod -- ls /app
kubectl exec -it nome-pod -c nome-contenitore -- bash
```

## kubectl apply
**Tags**: kubernetes, kubectl, apply, deploy, crea, manifest
**Keywords**: apply deploy crea manifest yaml configurazione applicare distribuire
**Short_Description**: Applica configurazione da file
**Long_Description**: Crea o aggiorna risorse Kubernetes da manifest YAML/JSON. Supporta file, URL e directory Kustomize con flag -k. Usa approccio dichiarativo - sicuro da eseguire più volte poiché applica solo le modifiche. Metodo preferito per workflow GitOps.

```sh
kubectl apply -f deployment.yaml
kubectl apply -f https://example.com/manifest.yaml
kubectl apply -k ./cartella-kustomize
```

## kubectl delete
**Tags**: kubernetes, kubectl, delete, rimuovi, cleanup
**Keywords**: delete rimuovi elimina cleanup distruggi risorsa cancellare pulire
**Short_Description**: Elimina risorse
**Long_Description**: Rimuove risorse Kubernetes per nome, file o selettore di label. Può eliminare risorse specifiche o tutte le risorse di un tipo usando --all. Supporta terminazione graduale con periodo di grazia configurabile. Usare con cautela in ambienti di produzione.

```sh
kubectl delete pod nome-pod
kubectl delete -f deployment.yaml
kubectl delete deployment nome-deployment
kubectl delete pods --all
```

## kubectl get services
**Tags**: kubernetes, kubectl, services, svc, networking, lista
**Keywords**: get services svc lista networking endpoint servizi rete
**Short_Description**: Lista servizi
**Long_Description**: Mostra i servizi Kubernetes che espongono applicazioni in esecuzione nei pod. Visualizza nome servizio, tipo (ClusterIP, NodePort, LoadBalancer), IP cluster, IP esterno, porte ed età. Usa -A per tutti i namespace o -o wide per dettagli aggiuntivi.

```sh
kubectl get services
kubectl get svc
kubectl get svc -A
kubectl get svc -o wide
```

## kubectl get deployments
**Tags**: kubernetes, kubectl, deployments, lista, apps
**Keywords**: get deployments lista applicazioni running stato distribuzioni
**Short_Description**: Lista deployment
**Long_Description**: Visualizza i deployment mostrando repliche desiderate vs pronte, stato aggiornamento, repliche disponibili ed età. I deployment gestiscono ReplicaSet che a loro volta gestiscono i pod. Essenziale per monitorare lo stato e la salute del deployment delle applicazioni.

```sh
kubectl get deployments
kubectl get deploy
kubectl get deployments -A
kubectl get deployments -o wide
```

## kubectl scale
**Tags**: kubernetes, kubectl, scale, repliche, orizzontale
**Keywords**: scale repliche orizzontale autoscale aumenta diminuisci scalare
**Short_Description**: Scala le repliche del deployment
**Long_Description**: Regola il numero di repliche di pod in un deployment, replicaset o statefulset. Aumenta o diminuisce le repliche per gestire i cambiamenti di carico. Può scalare da file o per nome risorsa. Le modifiche hanno effetto immediato ma creazione/terminazione pod segue lo scheduling di Kubernetes.

```sh
kubectl scale deployment nome-deployment --replicas=3
kubectl scale --replicas=5 -f deployment.yaml
```

## kubectl rollout
**Tags**: kubernetes, kubectl, rollout, deployment, aggiorna, rollback
**Keywords**: rollout deployment aggiorna rollback riavvia cronologia aggiornamento
**Short_Description**: Gestisci i rollout dei deployment
**Long_Description**: Controlla gli aggiornamenti dei deployment incluso monitoraggio stato, visualizzazione cronologia revisioni, rollback a versioni precedenti e riavvio deployment. Essenziale per gestire aggiornamenti applicazioni in sicurezza e recuperare da deployment falliti.

```sh
kubectl rollout status deployment/nome-deployment
kubectl rollout history deployment/nome-deployment
kubectl rollout undo deployment/nome-deployment
kubectl rollout restart deployment/nome-deployment
```

## kubectl port-forward
**Tags**: kubernetes, kubectl, port-forward, networking, debug
**Keywords**: port forward tunnel locale debugging accesso porta inoltro
**Short_Description**: Inoltra porta locale al pod
**Long_Description**: Crea un tunnel dalla macchina locale a un pod, servizio o deployment nel cluster. Mappa porta locale a porta container per accesso diretto senza esporre servizi pubblicamente. Utile per debug, test e accesso a risorse cluster localmente.

```sh
kubectl port-forward nome-pod 8080:80
kubectl port-forward svc/nome-servizio 8080:80
kubectl port-forward deployment/nome-deployment 8080:80
```

## kubectl create namespace
**Tags**: kubernetes, kubectl, namespace, crea, organizzazione
**Keywords**: crea namespace nuovo organizzazione isolamento creare
**Short_Description**: Crea un nuovo namespace
**Long_Description**: Crea un nuovo namespace per isolamento logico delle risorse nel cluster. I namespace forniscono ambito per nomi risorse, abilitano quote risorse e controllo accesso, e aiutano a organizzare cluster multi-tenant o multi-ambiente.

```sh
kubectl create namespace nome-namespace
kubectl create ns nome-namespace
```

## kubectl get nodes
**Tags**: kubernetes, kubectl, nodes, cluster, infrastruttura, lista
**Keywords**: get nodes lista cluster infrastruttura worker master nodi
**Short_Description**: Lista nodi del cluster
**Long_Description**: Visualizza i nodi del cluster mostrando stato (Ready/NotReady), ruoli (control-plane, worker), età e versione Kubernetes. Usa -o wide per dettagli aggiuntivi come IP interni/esterni, OS, kernel e runtime container. Critico per monitoraggio salute cluster.

```sh
kubectl get nodes
kubectl get nodes -o wide
kubectl describe node nome-nodo
```

## kubectl top
**Tags**: kubernetes, kubectl, metriche, risorse, cpu, memoria
**Keywords**: top metriche risorse cpu memoria utilizzo statistiche consumo
**Short_Description**: Mostra utilizzo risorse
**Long_Description**: Visualizza utilizzo corrente di CPU e memoria per nodi o pod. Richiede metrics-server installato nel cluster. Usa flag --containers per mostrare metriche container individuali dentro i pod. Essenziale per monitorare consumo risorse e identificare problemi performance.

```sh
kubectl top nodes
kubectl top pods
kubectl top pod nome-pod --containers
```

## kubectl config
**Tags**: kubernetes, kubectl, config, contesto, cluster, kubeconfig
**Keywords**: config contesto cluster kubeconfig switch corrente cambia configurazione
**Short_Description**: Gestisci kubeconfig
**Long_Description**: Gestisce configurazione kubectl inclusi contesti, cluster e utenti. Visualizza config corrente, lista contesti disponibili, cambia tra cluster/namespace e imposta namespace default. I contesti combinano informazioni cluster, utente e namespace per facile switching tra ambienti.

```sh
kubectl config view
kubectl config get-contexts
kubectl config use-context nome-contesto
kubectl config current-context
kubectl config set-context --current --namespace=nome-namespace
```
