# 🇬🇴 GoWatcher

**GoWatcher** est un projet pédagogique écrit en Go, développé dans le cadre d'un cours d'apprentissage de Go en **Master 2**.  
Il a pour objectif de familiariser les étudiants avec des concepts clés du langage Go : concurrence, gestion des erreurs, structure de projet, modularité, interfaces CLI (Cobra) et traitement de fichiers JSON.


## Description des branches 

Le projet a été refactorisé plusieurs fois afin d'introduire des nouveaux concepts au fur et à mesure. Ainsi les différentes branches se découpent comme tel :
* **`step1`:** Introduction des goroutines et les channels pour les tâches parallèles
* **`step2`:** Introduction des WaitGroups pour s'assurer que les goroutines se terminent correctement
* **`step3-errors`:** : Gestion des erreurs avancées et création d'erreurs personnalisées
* **`step4-subcommands`:** Intégration de Cobra pour construire des commandes
* **`main`:** qui correspond au **step5** avec l'ajout de la manipulation de données avec des fichiers JSON

## 📚 Objectifs pédagogiques

- Maîtriser la gestion de la **concurrence** en Go (`goroutines`, `WaitGroup`, `channels`)
- Structurer un projet modulaire en suivant les bonnes pratiques (`cmd/`, `internal/`, etc.)
- Créer une **application CLI** robuste avec `cobra`
- Manipuler des **fichiers JSON** pour configurer dynamiquement le comportement du programme
- Implémenter des **tests unitaires** (à venir)

## Fonctionnalités

- ✅ Vérifie l’accessibilité d’une liste d’URLs depuis un fichier `.json`
- 🛠 Affiche les statuts HTTP (succès ou erreur détaillée)
- ➕ Ajout dynamique de nouvelles URLs via une commande CLI
- 📁 Export des résultats dans un fichier JSON de rapport
- 🧵 Exécution concurrente des vérifications grâce à des `goroutines`
---

## Structure du projet

```
gowatcher/
├── cmd/                       # Commandes CLI
│   ├── add.go                 # Commande pour ajouter une URL au fichier JSON
│   ├── check.go               # Commande pour vérifier la disponibilité des URLs
│   └── root.go                # Commande racine qui initialise Cobra
├── internal/                  # Logique métier de l'application
│   ├── checker/
│   │   ├── check.go           # Vérifie l'accessibilité des URLs
│   │   └── errors.go          # Définit les erreurs personnalisées
│   ├── config/
│   │   └── config.go          # Chargement/sauvegarde du fichier JSON d'URLs
│   └── reporter/
│       └── report.go          # Exportation des résultats de vérification
├── reports/
│   └── results.json           # Fichier de résultats exporté (exemple)
├── urls.json                  # Fichier d'entrée avec la liste des URLs à surveiller
├── .gitignore                 # Fichiers/dossiers exclus du contrôle de version
├── go.mod                     # Dépendances du projet et nom du module
├── main.go                    # Point d'entrée de l'application
```

## Installation

Pour utiliser GoWatcher, assurez-vous d'avoir Go installé sur votre machine (version 1.18 ou supérieure recommandée).

1.  **Clonez le dépôt :**
    ```bash
    git clone https://github.com/axellelanca/gowatcher_correction.git)
    cd gowatcher_correction
    ```
2.  **Téléchargez les dépendances :**
    ```bash
    go mod tidy
    ```
3.  **Compilez l'exécutable :**
    ```bash
    go build -o gowatcher .
    ```
    Cela créera un exécutable nommé `gowatcher` (ou `gowatcher.exe` sur Windows) dans le répertoire courant.

---

##  Exemples d'utilisation
Voici comment utiliser les différentes commandes de GoWatcher.

### Vérifier des URLs
```bash
gowatcher check --input urls.json --output results.json
```

### Ajouter une nouvelle URL
```bash
gowatcher add --file urls.json --name "Site Test" --url "https://example.com" --owner "admin@example.com"
```


## Pistes d'amélioration 

* Ajout de tests unitaires et tests d’intégration
* Génération de rapports HTML
* Gestion de proxy ou en-têtes personnalisés
* Re-vérification périodique des URLs

## Licence
Projet open-source libre d'utilisation dans un cadre éducatif.
Développé dans un but exclusivement pédagogique.
