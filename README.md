# Projet Data


## Client : RATP

Ce dossier contient le code back-end et le code front-end du projet Data de l'équipe 1.

### Prérequis

L'installation de ce projet nécessite :  
- Go >1.13 (https://golang.org/dl/)  
- NPM (https://www.npmjs.com/get-npm) ou un autre package manager JS    

Si NPM est déjà installé, s'assurer d'avoir la dernière version :

```sudo npm install npm@latest -g```


## Installation serveur HTTP

Copier le dossier dans **$GOPATH/src/** puis :

```cd $GOPATH/src/data-eq1/API```

```make run``` ou ```go run main.go```  

Go va installer les packages nécessaires et servir l'application sur le port 80.  
Une documentation Swagger expliquant les différentes routes de l'API sera disponible à l'adresse suivante : http://localhost:80/api/v1/documentation/


## Installation serveur front-end

```cd /data-eq1/client```

Puis installer les dépendances : ```npm i```

Générer un build de production : ```npm run build```

Servir l'application sur le port 5000 : ```serve -s build```


## Parties incomplètes du projet

Add additional notes about how to deploy this on a live system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```


## Packages externes utilisées

### Go
- https://github.com/gorilla/mux : Multiplexer de requêtes HTTP (routing & dispatching) 
- https://github.com/lib/pq : Driver Postgres pour le package database/sql de la librairie standard
- https://github.com/paulmach/go.geojson : Encodeur et décodeur de GeoJSON

### React
- https://www.npmjs.com/package/mapbox-gl : Cartes interactives et personnalisables
- https://www.npmjs.com/package/react-apexcharts : Datavisualisations interactives
- https://www.npmjs.com/package/material-table : Tableaux personnalisables
- https://www.npmjs.com/package/react-router : Routeur côté client




