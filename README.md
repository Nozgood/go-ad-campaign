## Bonjour et bienvenue sur votre gestionnaire de campagnes 

Ici, vous pouvez créer, afficher, modifier et supprimer les campagnes publicitaires
Une campagne est composée :
- d'un nom
- d'une date de début,
- d'une date de fin,
- d'un prix (total en euros),
- d'un objectif d'affichage,
- et d'un prix par affichage (qui est calculé automatiquement)

### Comment accéder à l'application

#### Pour pouvoir y accéder grâce à docker hub : 
- lancez la commande `docker run -p 8080:8080 nozgood/goadcampaign`
Attention: il est nécéssaire que l'application soit rendu sur le port 8080 de votre machine (le port n'est pas géré grâce à une variable d'environnement)


#### Pour pouvoir y accéder grâce à Docker, depuis le repo cloné  : 
- Clonez le repo : `git clone https://github.com/Nozgood/go-ad-campaign.git`
- lancez d'abord la commange `docker build -t votrenomdimage .`
- lancez la commande `docker run -p 8080:8080 votrenomdimage`

- Si vous possédez déjà sur votre machine toutes les dépendances nécéssaires, depuis le repo cloné, à la racine : `go run main.go`

Pardonnez-moi d'avance pour le manque de style dont fait part cette application, ce n'était pas forcément l'objectif recherché sur cet exercice.
Le but étant que l'app soit fonctionnel, même si elle n'est pas belle :D 

Enjoy ! 