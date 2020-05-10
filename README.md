# Projet Data

## Client : RATP

Ce repo contient...

### Prérequis

L'installation de ce projet nécessite :  
- Go >1.13 (https://golang.org/dl/)  
- NPM (https://www.npmjs.com/get-npm) ou un autre package manager JS    

Si NPM est déjà installé, s'assurer d'avoir la dernière version :
```
sudo npm install npm@latest -g
```

## Installation serveur HTTP

Copier le dossier dans **$GOPATH/src/** puis :

```
cd $GOPATH/src/data-eq1/API
```
```make run``` ou ```go run main.go```  

Go va installer les packages nécessaires et servir l'API sur le port 80.  
Une documentation Swagger expliquant les différentes routes sera disponible à l'adresse suivante : http://localhost:80/api/v1/documentation/

## Installation serveur front-end

```
cd /data-eq1/client
```
Puis installer les dépendances
```
npm i
```
Générer un build de production
```
npm run build
```
Servir l'application sur le port 5000
```
serve -s build
```

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

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc