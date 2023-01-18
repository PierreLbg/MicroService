# Projet de cours de Microservices

Ce dépôt contient les projets du cours de microservices :

## Projet d'API REST

Ce projet est une API REST simple pour effectuer des opérations mathématiques telles que l'addition. Il est construit en utilisant un framework tel qu'Express.js. La route `/somme` permet d'additionner deux nombres en utilisant les paramètres `nombre1` et `nombre2` et renvoie le résultat sous la forme d'un nombre. Ce projet se trouve dans le répertoire `Rest-Somme`.

### Mise en route

Pour exécuter les projets, vous aurez besoin d'avoir Node.js, Express.js et pnpm installés sur votre machine.

1. Clonez le dépôt
2. Accédez au répertoire de projet `Rest-Somme`
3. Exécutez `pnpm install` pour installer les dépendances nécessaires
4. Pour lancer le serveur, utilisez `pnpm dev` dans un terminal.
5. Pour accéder au client, il suffit de se rendre à l'adresse http://localhost:9000/somme?nombre1=2&nombre2=4.

## Projet gRPC

Ce projet est un serveur gRPC construit en utilisant Node.js, express.js, la librairie grpc. Il permet de traiter les messages en utilisant la route `Traitement/Action` et de renvoyer une réponse. Le payload de la requête doit contenir un attribut "message". La réponse comprendra un nouvel attribut "AR" avec la valeur "OK". Ce projet se trouve dans le répertoire `grpc-js`.

### Mise en route

Pour exécuter les projets, vous aurez besoin d'avoir Node.js, pnpm et mosquitto installés sur votre machine. Il faut également installer les paquets @grpc/grpc-js et @grpc/proto-loader.

1. Clonez le dépôt
2. Accédez au répertoire de projet `grpc-js`
3. Exécutez `pnpm install` pour installer les dépendances nécessaires
4. Pour lancer le serveur gRPC, utilisez `pnpm start:server` dans un terminal, puis `pnpm start:client` dans un autre terminal.
5. Pour envoyer un message au serveur gRPC, utilisez la commande `mosquitto_pub -h 127.0.0.1  -t grpc_message -m "Ceci est un message"`,

## Projet Node-RED

Ce projet utilise Node-RED pour envoyer des messages au serveur gRPC en utilisant un subscriber MQTT pour récupérer les messages. Il est dépendant du projet gRPC et ne peut pas être exécuté sans celui-ci. Il se trouve dans le répertoire `node-red`.

### Mise en route

Pour exécuter les projets, vous aurez besoin d'avoir Node-RED, Node.js, pnpm et mosquitto installés sur votre machine. 

1. Clonez le dépôt
2. Accédez au répertoire de projet `node-red`
3. Exécutez `pnpm install` pour installer les dépendances nécessaires
4. Ouvrez Node-RED et importez le fichier Tntégration.json situé dans le répertoire `node-red`
5. Éffectuer la mise en route du Projet gRPC en allant jusqu'à la requête `pnpm start:server`
6. Vous pouvez maintenant utilisez la commande `mosquitto_pub -h 127.0.0.1  -t node_red -m "Ceci est un message"`,