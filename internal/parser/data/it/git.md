# Comandi Git

## git pull
**Tags**: git, version-control, sync, remote, fetch
**Keywords**: pull fetch remote update sync download scaricare aggiornare sincronizzare
**Short_Description**: Scarica e unisce modifiche dal repository remoto
**Long_Description**: Scarica commit, file e riferimenti da un repository remoto e li integra immediatamente nel branch corrente. È essenzialmente una combinazione di git fetch seguito da git merge. Utile per sincronizzare il lavoro locale con quello del team.

```sh
git pull origin main
```

## git push
**Tags**: git, version-control, sync, remote, upload
**Keywords**: push remote upload sync send publish caricare pubblicare inviare sincronizzare
**Short_Description**: Carica i commit locali sul repository remoto
**Long_Description**: Invia i commit locali al repository remoto per condividerli con il team. Il flag --force-with-lease è più sicuro di --force perché fallisce se qualcun altro ha pushato nel frattempo. Essenziale per pubblicare le modifiche e collaborare.

```sh
git push origin main
git push --force-with-lease
```

## git stash
**Tags**: git, version-control, temporary, save, workspace
**Keywords**: stash save temporary work-in-progress wip store salvare temporaneo accantonare
**Short_Description**: Salva temporaneamente file modificati
**Long_Description**: Accantona temporaneamente le modifiche nella directory di lavoro per poter cambiare branch senza committare. Usa pop per riapplicare le modifiche, list per vedere tutti gli stash, e apply per riapplicare uno stash specifico mantenendolo nella lista.

```sh
git stash
git stash pop
git stash list
git stash apply stash@{0}
```

## git commit
**Tags**: git, version-control, save, snapshot
**Keywords**: commit save snapshot message record changes salvare registrare modifiche
**Short_Description**: Registra le modifiche nel repository
**Long_Description**: Crea uno snapshot delle modifiche staged con un messaggio descrittivo. Il flag -m specifica il messaggio inline, --amend modifica l'ultimo commit, --no-edit mantiene il messaggio precedente. Ogni commit rappresenta un punto nella storia del progetto.

```sh
git commit -m "Add new feature"
git commit --amend
git commit --amend --no-edit
```

## merge and delete branch
**Tags**: git, github, branch, merge, delete, pull-request, pr
**Keywords**: merge branch delete remove pull-request pr gh cleanup unire eliminare pulizia
**Short_Description**: Unisce pull request ed elimina il branch automaticamente
**Long_Description**: Unisce una pull request e elimina il branch associato in un singolo comando usando GitHub CLI. Utile per mantenere pulito il repository rimuovendo branch dopo il merge. Può essere fatto anche manualmente con git branch -d per branch locali e git push --delete per branch remoti.

```sh
gh pr merge 38 --merge --delete-branch
git branch -d feature-branch
git push origin --delete feature-branch
```

## create branch
**Tags**: git, branch, create, new
**Keywords**: branch create new checkout switch creare nuovo ramo
**Short_Description**: Crea e passa a un nuovo branch
**Long_Description**: Crea un nuovo branch e ci si sposta in un singolo comando. git checkout -b è il metodo tradizionale, git switch -c è il comando più moderno. Essenziale per lavorare su feature o fix isolati senza modificare il branch principale.

```sh
git checkout -b feature/new-feature
git switch -c feature/new-feature
```

## git rebase
**Tags**: git, rebase, history, clean
**Keywords**: rebase interactive squash history rewrite riscrivere cronologia pulire
**Short_Description**: Riorganizza i commit per una cronologia più pulita
**Long_Description**: Riapplica i commit sopra un altro punto base per linearizzare la cronologia. Rebase interattivo (-i) permette di squashare, riordinare, o modificare commit. Usa --continue dopo aver risolto conflitti. Attenzione: non fare rebase di commit già pushati pubblicamente.

```sh
git rebase main
git rebase -i HEAD~3
git rebase --continue
```

## git reset
**Tags**: git, reset, undo, rollback
**Keywords**: reset undo rollback uncommit restore annullare ripristinare
**Short_Description**: Resetta commit o modifiche
**Long_Description**: Sposta HEAD e opzionalmente l'indice e la directory di lavoro. --soft mantiene le modifiche staged, --mixed (default) unstage le modifiche, --hard scarta tutto. Utile per annullare commit locali o ripulire la working directory. Attenzione con --hard: le modifiche vengono perse.

```sh
git reset --soft HEAD~1
git reset --hard HEAD~1
git reset --mixed HEAD~1
```
