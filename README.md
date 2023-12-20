# WebToolKit

Ceci est un modèle d'application web multi-tenants en Go pour la partie serveur et Svelte pour la partie cliente.

Le multi-tenant est assuré par le host dans l'url d'accès au service. Par exemple :

`http://realm.local.tld`

utilisera le tenant "realm" dans le reste de l'application (login, handlers, ...)

Pour tester, il faut utiliser http://test.localhost:8000 (le tenant 'test' est défini dans server.go)

La partie css est gérée par Tailwindcss.

## Pré-requis

- Go : pour gérer le code en Go

- npm : pour gérer le code en Javascript pour Svelte

- Docker et Docker Compose : pour "packager" l'application et en faciliter le déploiement. Peut aussi éviter d'installer Go car le Dockerfile contient ce qu'il faut pour compiler du Go via un container dédié.

- [Watcher](https://github.com/sipkg/watcher) : recompile le code Go dès qu'un fichier est modifié.

- [GoConvey](https://github.com/smartystreets/goconvey) : relance les tests et affiche le résultat dans le navigateur à chaque modification.

## Installation

```
make install
```

Installe les dépendances Go (backend) et Javascript (frontend).

## Utilisation

```
make
```

Lance dans 3 panels tmux :

- GoConvey pour lancer automatiquement les tests Go
- Watcher pour compiler et lancer le backend 
- Svelte via 'npm' pour bundler et lancer le frontend

A chaque modification, le composant correspondant (front ou back) est recompilé et relancé.

L'architecture côté serveur respecte les principes de séparation des responsabilités :

- 'biz' est en charge de la logique métier et de la structuration des données. C'est ici qu'on devrait trouver ce qui fait la particularité de l'application.

- 'ui' contient les interfaces utilisateurs (au sens large) :
    * 'api' pour les interfaces http
    * 'web' pour l'interface web humain-machine (client svelte)

- 'cmd' contient les commandes de l'application. A minima, il contient le serveur http dans le dossier "server". 

- 'ext' regroupe les éléments externes à l'application, comme le stockage des données.

## Stratégie d'authentification avec jeton (token)

Pour gérer l'authentification des utilisateurs, on leur demande un identifiant et un mot de passe, dont on contrôle la validité. Si les données sont correctes, on génère un jeton signé et crypté qui contient la structure suivante :

```
identifiant|adresse IP|horodatage
```

Le jeton est envoyé au navigateur de l'utilisateur sous la forme d'un 'cookie' : le navigateur va donc nous renvoyer ce 'cookie' à chaque requête. On pourra alors contrôler l'accès légitime de cette requête :

1. on décrypte le contenu du 'cookie' : si on y arrive c'est que la signature est bonne.
2. on vérifie que l'adresse IP est toujours celle de connexion
3. on vérifie que cela ne fait pas trop longtemps que la connexion a eu lieu

## Test de l'API avec curl

```sh
# On s'authentifie en envoyant un JSON et on sauvegarde le cookie dans un fichier
curl -v -c /tmp/cookie.txt -d '{"username":"test","password":"test"}' http://test.localhost:8000/login
# On peut appeler un service en réutilisant le fichier cookie
curl -v -b /tmp/cookie.txt http://test.localhost:8000/hello\?nom\=le%20monde
# On se déconnecte en modifiant le fichier cookie
curl -v -c /tmp/cookie.txt http://test.localhost:8000/logout
```
