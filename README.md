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

### Fausses donneés

Les données suivantes ont été simulées ou *randomisées* afin de rendre la présentation au client plus réaliste. Ces données doivent provenir de l'application mobile (solution finale prototypée).

- Onglet **Insights** : données en haut de page et en bas de page, indices du graphique de droite (faux indices, il s'agit des unités de chaque critère)

- Onglet **Voyageurs** : données en haut de page, les noms et prénoms des utilisateurs (données MaRATP non nominatives), les statistiques/badges/la relation des utilisateurs.
Les données présentes dans la table *trajets* de la base de données correspondent à des exemples de trajets générés par les utilisateurs de l'app mobile prototypée. Ces trajets ont été assignés à des clients par leur identifiant (table fournie en début de projet). La barre de recherche d'un utilisateur ne fonctionne avec qu'avec cet identifiant, pour les clients ayant effectué des trajets.

### Parties incomplètes

- La page de connexion n'authentifie pas réellement l'utilisateur
- Page **Données** : La météo affichée pour les graphiques de la qualité de l'air est la même pour tout Paris, elle ne change pas pour chaque arrondissement. Le bouton d'ajout d'une table n'en crée pas réellement une. Le bouton de modification d'une table permet simplement de la visualiser, la fonctionnalité d'ajout et de modification des entrées ne fonctionne que côté client. Le bouton de téléchargement de la table au format CSV est décoratif.

### Améliorations souhaitées

Responsive, loaders, memory leaks, routing

Lorem ipsum lorem ipsum lorem ipsum lorem ipsum

```
Exemple
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




