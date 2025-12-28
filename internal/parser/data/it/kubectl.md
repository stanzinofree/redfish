# Scorciatoie e Suggerimenti Kubectl

## kubectl aliases
**Tags**: kubectl, alias, scorciatoie, produttività
**Keywords**: alias scorciatoie veloci produttività suggerimenti rapidi
**Short_Description**: Alias comuni per kubectl
**Long_Description**: Configura alias di shell per ridurre la digitazione e velocizzare il workflow kubectl. Questi alias creano scorciatoie per comandi kubectl usati frequentemente. Aggiungili al tuo .bashrc o .zshrc per disponibilità permanente. Migliora significativamente la produttività per operazioni Kubernetes quotidiane.

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
**Keywords**: completion autocompletamento bash zsh shell setup configurazione
**Short_Description**: Configura autocompletamento shell
**Long_Description**: Abilita autocompletamento tab per comandi kubectl, flag e nomi risorse nella tua shell. Supporta bash, zsh e altre shell. Velocizza significativamente l'inserimento comandi e riduce errori auto-completando nomi risorse e flag. Deve essere configurato una volta per shell.

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
**Keywords**: explain documentazione api riferimento aiuto campi spiegazione
**Short_Description**: Ottieni documentazione per le risorse
**Long_Description**: Visualizza documentazione per campi risorse Kubernetes direttamente dallo schema API. Naviga campi annidati con notazione punto. Essenziale per comprendere campi disponibili, loro tipi e requisiti quando scrivi manifest. Funziona offline usando schema locale.

```sh
kubectl explain pods
kubectl explain pods.spec
kubectl explain pods.spec.containers
kubectl explain deployment.spec.replicas
```

## kubectl diff
**Tags**: kubectl, diff, confronta, modifiche, anteprima
**Keywords**: diff confronta modifiche anteprima dry-run prova differenze
**Short_Description**: Anteprima modifiche prima dell'applicazione
**Long_Description**: Mostra quali modifiche verrebbero apportate da kubectl apply prima di applicarle effettivamente. Confronta manifest locali con stato cluster live e visualizza differenze. Essenziale per deployment sicuri - rivedi modifiche prima che influenzino produzione. Funziona con file e directory Kustomize.

```sh
kubectl diff -f deployment.yaml
kubectl diff -k ./cartella-kustomize
```

## kubectl wait
**Tags**: kubectl, wait, condizione, ready, sincronizzazione
**Keywords**: wait aspetta condizione ready sincronizzazione timeout attesa
**Short_Description**: Attendi una condizione
**Long_Description**: Blocca fino al raggiungimento di una condizione specificata su risorse Kubernetes. Utile in script e pipeline CI/CD per assicurare che le risorse siano pronte prima di procedere. Supporta condizioni come ready, available, deleted con timeout configurabili. Restituisce codice uscita al timeout.

```sh
kubectl wait --for=condition=ready pod/nome-pod
kubectl wait --for=condition=available deployment/nome-deployment --timeout=60s
kubectl wait --for=delete pod/nome-pod --timeout=30s
```

## kubectl run quick pod
**Tags**: kubectl, run, pod, veloce, test, debug
**Keywords**: run pod veloce test debug temporaneo effimero eseguire
**Short_Description**: Esegui rapidamente un pod di test
**Long_Description**: Crea un pod effimero per test rapidi o debug. Il flag --rm elimina il pod all'uscita, -it fornisce terminale interattivo. Utile per testare connettività rete, eseguire comandi rapidi o debuggare problemi cluster senza creare risorse permanenti.

```sh
kubectl run test --image=nginx --rm -it -- bash
kubectl run debug --image=busybox --rm -it -- sh
kubectl run curl --image=curlimages/curl --rm -it -- sh
```

## kubectl patch
**Tags**: kubectl, patch, aggiorna, modifica, inline
**Keywords**: patch aggiorna modifica inline veloce cambia modificare
**Short_Description**: Modifica risorse inline
**Long_Description**: Aggiorna campi specifici di risorse usando patch JSON/YAML senza bisogno di file manifest completi. Utile per modifiche rapide a risorse in esecuzione. Supporta strategic merge patch, JSON merge patch e JSON patch. Più veloce di edit per aggiornamenti scriptati.

```sh
kubectl patch deployment nome-deployment -p '{"spec":{"replicas":3}}'
kubectl patch pod nome-pod -p '{"metadata":{"labels":{"env":"prod"}}}'
```

## kubectl label
**Tags**: kubectl, label, metadata, organizza, seleziona
**Keywords**: label etichetta metadata organizza seleziona tag etichettare
**Short_Description**: Aggiungi o modifica etichette
**Long_Description**: Gestisce etichette su risorse Kubernetes. Le etichette sono coppie chiave-valore usate per organizzare, selezionare e filtrare risorse. Aggiungi etichette con chiave=valore, rimuovi con chiave-. Può applicare a più risorse usando --all o selettori. Critico per service discovery e deployment.

```sh
kubectl label pods nome-pod env=prod
kubectl label pods nome-pod env-
kubectl label pods --all env=staging
```

## kubectl annotate
**Tags**: kubectl, annotate, metadata, documentazione
**Keywords**: annotate annotazione metadata documentazione note annotare
**Short_Description**: Aggiungi o modifica annotazioni
**Long_Description**: Gestisce annotazioni su risorse. Le annotazioni memorizzano metadata non identificativi come info build, note deployment o dati tool-specifici. A differenza delle etichette, non vengono usate per selezione. Rimuovi annotazioni con chiave-. Utile per documentazione e integrazione con tool esterni.

```sh
kubectl annotate pods nome-pod description="pod di test"
kubectl annotate pods nome-pod description-
```

## kubectl drain node
**Tags**: kubectl, drain, nodo, manutenzione, evict
**Keywords**: drain svuota nodo manutenzione evict cordon uncordon evacuare
**Short_Description**: Svuota nodo per manutenzione
**Long_Description**: Evacua in sicurezza i pod da un nodo per manutenzione. Cordon previene scheduling nuovi pod, drain evacua pod esistenti rispettando PodDisruptionBudgets. Usa --ignore-daemonsets per ignorare pod DaemonSet. Uncordon marca il nodo schedulabile di nuovo dopo manutenzione.

```sh
kubectl drain nome-nodo --ignore-daemonsets
kubectl cordon nome-nodo
kubectl uncordon nome-nodo
```

## kubectl debug
**Tags**: kubectl, debug, troubleshoot, effimero, contenitore
**Keywords**: debug troubleshoot diagnostica effimero contenitore diagnosi debuggare
**Short_Description**: Debug pod in esecuzione
**Long_Description**: Crea container debug effimeri in pod o nodi in esecuzione senza modificare la specifica pod. Utile per debuggare immagini distroless o minimali che mancano di strumenti debug. Può attaccarsi a container esistenti o creare nuovi container debug con toolset completi.

```sh
kubectl debug nome-pod -it --image=busybox
kubectl debug nome-pod -it --image=ubuntu --target=nome-contenitore
kubectl debug node/nome-nodo -it --image=ubuntu
```

## kubectl events
**Tags**: kubernetes, kubectl, eventi, logs, cluster, troubleshoot
**Keywords**: eventi logs cluster troubleshoot cronologia diagnostica eventi
**Short_Description**: Visualizza eventi del cluster
**Long_Description**: Mostra eventi cluster-wide che mostrano cambiamenti di stato, errori e avvisi. Gli eventi sono limitati nel tempo (solitamente 1 ora) e forniscono insight sulle operazioni cluster. Ordina per timestamp, filtra per namespace o tipo. Essenziale per troubleshooting fallimenti pod e problemi cluster.

```sh
kubectl get events
kubectl get events --sort-by=.metadata.creationTimestamp
kubectl get events -n nome-namespace
kubectl get events --field-selector type=Warning
```
