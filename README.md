# Projet de cours de Microservices

Ce dépôt contient les projets du cours de microservices :

### Sommaire
- [Prérequis](#prérequis)
- [Mise en route générale](#mise-en-route-générale)
- [Projet d'API REST](#projet-dapi-rest)
- [Projet gRPC](#projet-grpc)
- [Projet Node-RED](#projet-node-red)
- [Projet Node-red_Multiservice](#projet-node-red_multiservice)
- [Projet Go-kit](#projet-go-kit)

### Prérequis

Avant de lancer les projets, vous aurez besoin d'installer les logiciels et librairies suivants sur votre machine :
- Node.js
- pnpm
- mosquitto
- Node-RED
- Go
- Go-kit


### Mise en route générale

1. Clonez le dépôt
2. Accédez au répertoire de projet correspondant (`Rest-Somme`, `grpc-js`, `node-red`, `node-red_multiservice`, `go-kit`)
3. Exécutez `pnpm install` pour installer les dépendances nécessaires

## Projet d'API REST

Ce projet est une API REST simple pour effectuer des opérations mathématiques telles que l'addition. Il est construit en utilisant un framework tel qu'Express.js. La route `/somme` permet d'additionner deux nombres en utilisant les paramètres `nombre1` et `nombre2` et renvoie le résultat sous la forme d'un nombre. 
> Ce projet se trouve dans le répertoire `Rest-Somme`.

### Mise en route

1. Pour lancer le serveur, utilisez `pnpm dev` dans un terminal.
2. Pour accéder au client, utilisez la commande suivante dans un terminal : 
    curl --request GET \
    --url 'http://localhost:9000/somme?nombre1=1&nombre2=3'


## Projet gRPC

Ce projet est un serveur gRPC construit en utilisant la librairie grpc de Node.js. Il permet de traiter les messages en utilisant la route `Traitement/Action` et de renvoyer une réponse. Le payload de la requête doit contenir un attribut "message". Les messages sont récupérés à partir d'un subscriber MQTT. La réponse comprendra un nouvel attribut "AR" avec la valeur "OK". 
> Ce projet se trouve dans le répertoire `grpc-js`.

### Mise en route

1. Pour lancer le serveur gRPC, utilisez `pnpm start:server` dans un terminal, puis `pnpm start:client` dans un autre terminal.
2. Pour envoyer un message au serveur gRPC, utilisez la commande `pnpm start:message`


## Projet Node-RED

Ce projet utilise Node-RED pour envoyer des messages au serveur gRPC en utilisant un subscriber MQTT pour récupérer les messages. Il est dépendant du projet gRPC et ne peut pas être exécuté sans celui-ci. C'est le serveur du projet gRPC qui doit tourné (`grpc-js $pnpm start:server`)
> Ce projet se trouve dans le répertoire `node-red`.

### Mise en route

1. Ouvrez Node-RED en utilisant la commande `pnpm start:node_red`
2. Importez le fichier Tntégration.json situé dans le répertoire `node-red`
3. Vous pouvez maintenant utilisez la commande `pnpm start:message` pour envoyer des messages via Node-RED


## Projet Node-red_Multiservice

Ce projet utilise Node-RED pour envoyer des requêtes REST à un service, qui va ensuite requêter un autre service, cachant ainsi l'existence de ce dernier à l'utilisateur final. Il se trouve dans le répertoire `node-red_multiservice`.

### Mise en route

1. Exécutez la commande `pnpm start:service1` pour lancer le premier service et accéder à Node-RED sur le port 1882
2. Exécutez la commande `pnpm start:service2` pour lancer le second service et accéder à Node-RED sur le port 1881
3. Utilisez la commande dans un terminal : 
    curl --request GET \
    url 'http://localhost:1882/s1?type=OK'` pour envoyer une requête au service
   ou utilisez un logiciel de requêtage pour envoyer une requête à l'URL 'http://localhost:1882/s1?type=OK'


## Projet Go-kit

Ce projet utilise Go-kit pour créer des microservices en utilisant le pattern "Kit". Il comprend deux parties : "example" et "exercice".

### Partie example

La partie example contient des exemples tirés directement de la documentation de Go-kit. Ces exemples permettent de comprendre les différentes fonctionnalités de Go-kit.

### Partie exercice

La partie exercice contient deux exercices. Le premier, dans le dossier "exercice1", contient un seul fichier "main.go" qui permet d'effectuer une somme comme pour le projet "Rest-somme".

Le deuxième exercice, dans le dossier "exercice2", reprend le même exercice que le premier en séparant le fichier "main.go" en différents fichiers correspondant aux couches de service. Il ajoute également un logger pour suivre les événements.

### Mise en route

Pour exécuter les projets, vous aurez besoin d'avoir Go installé sur votre machine.

1. Accédez au répertoire de projet "go-kit/exercice/"
2. Pour lancer l'exercice 1 utilisez la commande "go run exercice1/main.go"
3. Pour lancer l'exercice 2 utilisez la commande "go run exercice2/*.go"
