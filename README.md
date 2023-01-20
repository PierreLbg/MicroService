# Projet de cours de Microservices

Ce dépôt contient les projets du cours de microservices :

## Prérequis

Avant de lancer les projets, vous aurez besoin d'installer les logiciels et librairies suivants sur votre machine :
- Node.js
- pnpm
- mosquitto
- Node-RED

## Mise en route générale
1. Clonez le dépôt
2. Accédez au répertoire de projet correspondant (`Rest-Somme`, `grpc-js`, `node-red`)
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
