# Comandi Docker

## docker ps
**Tags**: docker, container, list, running, status
**Keywords**: ps list containers running active status show elenco lista contenitori attivi
**Short_Description**: Elenca i container in esecuzione
**Long_Description**: Mostra informazioni sui container Docker in esecuzione inclusi ID, immagine, comando, data creazione, stato, porte e nomi. Usa il flag -a per includere anche i container fermati. Utile per monitorare lo stato dei container e ottenere gli ID per altri comandi.

```sh
docker ps
docker ps -a
docker ps --format "table {{.Names}}\t{{.Status}}"
```

## docker build
**Tags**: docker, image, build, create, dockerfile
**Keywords**: build image create dockerfile compile costruire immagine creare compilare
**Short_Description**: Costruisce un'immagine Docker da Dockerfile
**Long_Description**: Crea un'immagine Docker da un Dockerfile e un contesto di build (solitamente la directory corrente). Il flag -t assegna un tag all'immagine con nome e versione opzionale. Usa --no-cache per ricostruire da zero senza usare layer in cache. Le build multi-stage possono targetizzare stadi specifici con --target.

```sh
docker build -t myapp:latest .
docker build --no-cache -t myapp:latest .
docker build --target production -t myapp:prod .
```

## docker run
**Tags**: docker, container, run, start, execute
**Keywords**: run start execute container create launch eseguire avviare creare lanciare
**Short_Description**: Esegue un comando in un nuovo container
**Long_Description**: Crea e avvia un nuovo container da un'immagine. Il flag -d esegue in modalità detached (background), -p mappa le porte (host:container), -it fornisce terminale interattivo, -v monta volumi, e --rm rimuove il container dopo l'uscita. Comando essenziale per eseguire applicazioni containerizzate.

```sh
docker run -d -p 8080:80 nginx
docker run -it ubuntu bash
docker run --rm -v $(pwd):/app myapp
```

## docker compose up
**Tags**: docker, compose, orchestration, multi-container, start
**Keywords**: compose up start orchestration multi-container services avviare servizi orchestrazione
**Short_Description**: Avvia applicazione multi-container
**Long_Description**: Legge docker-compose.yml e avvia tutti i servizi definiti. Il flag -d esegue in modalità detached, --build forza la ricostruzione delle immagini prima dell'avvio, e --scale permette di eseguire istanze multiple di un servizio. Crea reti e volumi come definiti nel file compose.

```sh
docker compose up -d
docker compose up --build
docker compose up --scale web=3
```

## docker compose down
**Tags**: docker, compose, stop, cleanup, remove
**Keywords**: compose down stop remove cleanup teardown fermare rimuovere pulizia
**Short_Description**: Ferma e rimuove container e reti
**Long_Description**: Ferma tutti i servizi e rimuove container e reti creati da docker compose up. Il flag -v rimuove anche i volumi nominati, e --rmi all rimuove tutte le immagini usate dai servizi. Lascia i volumi intatti per default per preservare i dati.

```sh
docker compose down
docker compose down -v
docker compose down --rmi all
```

## docker logs
**Tags**: docker, logs, debugging, output
**Keywords**: logs debug output tail follow registro log debug
**Short_Description**: Visualizza i log del container
**Long_Description**: Recupera e mostra i log da un container in esecuzione o fermato. Il flag -f segue l'output dei log in tempo reale (come tail -f), --tail N mostra solo le ultime N righe, e --since può filtrare i log per timestamp. Essenziale per il debug di applicazioni containerizzate.

```sh
docker logs container-name
docker logs -f container-name
docker logs --tail 100 container-name
```

## docker exec
**Tags**: docker, exec, shell, command, debug
**Keywords**: exec shell bash command execute debug eseguire comando terminale
**Short_Description**: Esegue un comando in un container in esecuzione
**Long_Description**: Esegue un nuovo comando in un container già in esecuzione. I flag -it forniscono un terminale interattivo, comunemente usato per aprire una shell (bash, sh) per il debug. Può anche eseguire singoli comandi senza modalità interattiva. Il container deve essere in esecuzione.

```sh
docker exec -it container-name bash
docker exec container-name ls /app
```

## docker prune
**Tags**: docker, cleanup, prune, disk-space
**Keywords**: prune cleanup remove disk space garbage pulizia rimuovere spazio disco
**Short_Description**: Rimuove oggetti Docker inutilizzati
**Long_Description**: Pulisce risorse Docker inutilizzate per liberare spazio su disco. System prune rimuove container fermati, reti inutilizzate, immagini dangling e cache di build. Il flag -a rimuove anche le immagini inutilizzate (non solo dangling). Volume e image prune targetizzano tipi di risorse specifici.

```sh
docker system prune
docker system prune -a
docker volume prune
docker image prune -a
```
