# Projet Data

## Client : RATP

Ce repository contient le code back-end et le code front-end du projet Data de l'équipe 1 : le dashboard permettant de piloter la solution finale (mobile) **Citi+**.  

### Prérequis

L'installation de ce projet nécessite :  
- Go > 1.13 
- NPM 

## Installation serveur HTTP

Cloner ce repository dans **$GOPATH/src/**.  
Se rendre dans **/API/model/**, puis ouvrir les fichiers aux emplacements suivants :  
- /airparif/APIKey.go
- /elevation/APIKey.go
- /paris/APIKey.go  

Remplacer les **APIKey** par les clés API nécessaires (cf. mail).  
Ouvrir ./database/database.go et remplacer *username:password@host:port/databasename* par les bons identifiants.  

Installer les packages nécessaires et servir l'API sur le port 80 :  

```
cd API
go get
go run main.go
```

La route suivante permet d'accéder à la documentation Swagger expliquant les routes et d'y effectuer des tests : http://localhost:80/api/v1/documentation/

## Installation serveur front-end

Se rendre dans le dossier **client** puis ouvrir le fichier à l'emplacement suivant : src/components/Map/index.js  

Remplacer **MapboxToken** à la ligne 44 par la bonne clé API (cf. mail).  

Installer les dépendances, créer un build de production et le servir sur le port 5000 :  
```
cd client
npm i
npm run build
serve -s build
```

## Modèles de compréhension

Les données suivantes ont été simulées ou *randomisées* afin de rendre la présentation au client réaliste. Ces données sont censées provenir de l'application mobile (solution finale prototypée).

- Onglet **Insights** : données en haut de page et en bas de page, indices du graphique de droite (faux indices, il s'agit des unités de chaque critère).

- Onglet **Voyageurs** : données en haut de page, les noms et prénoms des utilisateurs (données MaRATP non nominatives), les statistiques/badges/la relation des utilisateurs.

Les données présentes dans la table *trajets* de la base de données correspondent à des exemples de trajets générés par les utilisateurs de l'app mobile prototypée. Ces trajets ont été assignés à des clients par leur identifiant (table fournie en début de projet). La barre de recherche d'un utilisateur (onglet **Voyageurs**) ne fonctionne avec qu'avec cet identifiant, pour les clients ayant effectué des trajets.

## Parties incomplètes

- La page de connexion n'authentifie pas d'utilisateur.
- Page **Données** : La météo affichée pour les graphiques de la qualité de l'air est la même pour tout Paris, elle ne change pas pour chaque arrondissement. Le menu de changement de *timeframes* n'est pas fonctionnel.  
Le bouton d'ajout d'une table n'en crée pas réellement une. Le bouton de modification d'une table permet simplement de la visualiser, la fonctionnalité d'ajout et de modification des entrées ne fonctionne que côté client. Le bouton de téléchargement d'une table au format CSV n'est pas fonctionnel.

## Packages externes utilisés

### Go
- https://github.com/gorilla/mux : Multiplexer de requêtes HTTP (routing & dispatching) 
- https://github.com/lib/pq : Driver Postgres pour le package database/sql de la librairie standard
- https://github.com/paulmach/go.geojson : Encodeur et décodeur d'objets GeoJSON

### React
- https://www.npmjs.com/package/mapbox-gl : Cartes interactives et personnalisables
- https://www.npmjs.com/package/react-apexcharts : Datavisualisations interactives
- https://www.npmjs.com/package/material-table : Tableaux personnalisables et 
- https://www.npmjs.com/package/react-router : Routeur côté client




