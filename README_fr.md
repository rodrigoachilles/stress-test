Traductions:

* [Anglais](README.md)
* [Portugais (Brésil)](README_pt_br.md)

---

# 🌐 CLI de Test de Stress (stress-test)

![Project Logo](assets/stress_test-logo.png)

Bienvenue dans le CLI de test de stress développé en Go ! Ce projet vous permet d'effectuer des tests de stress sur un service web en spécifiant l'URL, le nombre total de requêtes et le nombre d'appels simultanés.

## 📑&nbsp;Table des Matières

- [📖 Introduction](#introduction)
- [🛠 Prérequis](#prérequis)
- [⚙️ Installation](#installation)
- [🚀 Utilisation](#utilisation)
- [🔍 Exemples](#exemples)
- [🤝 Contribution](#contribution)
- [📜 Licence](#licence)

## 📖&nbsp;Introduction

Cet outil CLI vous permet d'effectuer des tests de stress sur un service web en spécifiant des paramètres tels que l'URL, le nombre total de requêtes et le nombre d'appels simultanés. Il génère un rapport détaillé avec des informations sur les requêtes effectuées.

## 🛠&nbsp;Prérequis

Assurez-vous d'avoir les éléments suivants installés avant de continuer :

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)

## ⚙️&nbsp;Installation

1. Clonez ce dépôt :

    ```sh
    git clone git@github.com:rodrigoachilles/stress-test.git
    cd stress-test
    ```

2. Construisez l'image Docker :

    ```sh
    docker build -t stress-test .
    ```

## 🚀&nbsp;Utilisation

Après avoir construit l'image Docker, vous pouvez exécuter le test de stress en utilisant la commande suivante :

```sh
docker run stress-test --url=http://example.com --requests=1000 --concurrency=10
```

### 📄&nbsp;Paramètres de la Ligne de Commande

- `--url` : URL du service web à tester.
- `--requests` : Nombre total de requêtes à effectuer.
- `--concurrency` : Nombre d'appels simultanés.

### 📋&nbsp;Rapport

Après le test, un rapport sera généré avec les informations suivantes :
- Temps total écoulé
- Nombre total de requêtes effectuées
- Nombre de requêtes avec le statut HTTP 200
- Répartition des autres codes de statut HTTP (par exemple, 404, 500, etc.)

## 🔍&nbsp;Exemples

Voici quelques exemples d'utilisation du CLI de test de stress :

- Tester un service web avec 1000 requêtes et 10 appels simultanés :
    ```sh
    docker run stress-test --url=http://example.com --requests=1000 --concurrency=10
    ```

## 🤝&nbsp;Contribution

N'hésitez pas à ouvrir des issues ou à soumettre des "pull requests" pour des améliorations et des corrections de bugs.

## 📜&nbsp;Licence

Ce projet est sous licence MIT.
