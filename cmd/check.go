package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/axellelanca/gowatcher_correction/internal/checker"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Vérifie l'accessibilité d'une liste d'URLs.",
	Long:  `La commande 'check' parcourt une liste prédéfinie d'URLs et affiche leur statut d'accessibilité en utilisant des goroutines pour la concurrence.`,
	Run: func(cmd *cobra.Command, args []string) { // La fonction `Run` est le cœur de la sous-commande.
		// Elle est exécutée lorsque l'utilisateur tape `gowatcher check`.
		// `cmd` représente la commande elle-même, `args` sont les arguments positionnels passés après la commande.

		var targets = []string{
			"https://www.google.com",
			"https://www.notarealwebsite.abc",
			"https://github.com",
			"https://www.movie.database/film/details",
			"https://www.gaming.news/release/new-game",
			"https://www.health.clinic/appointment/online",
			"https://www.car.manufacturer/model/electric",
			"https://www.home.decor/ideas/living-room",
			"https://www.environmental.org/project/clean-water",
			"https://www.space.agency/mission/mars",
			"https://www.fashion.magazine/trend/summer",
			"https://www.tech.conference/schedule/day1",
			"https://www.food.blog/recipe/dessert",
			"https://www.online.course/programming/python",
			"https://www.travel.guide/city/paris",
			"https://www.music.label/artist/new-album",
			"https://www.sports.club/events/match",
			"https://www.photography.tips/technique/lighting",
			"https://www.diy.tools/review/drill",
			"https://www.pet.vet/service/vaccination",
			"https://www.gardening.store/seeds/flower",
			"https://www.finance.advice/retirement/planning",
			"https://www.history.podcast/episode/ww2",
			"https://www.language.exchange/partner/find",
			"https://www.book.review/author/classic",
			"https://www.movie.review/genre/comedy",
			"https://www.gaming.forum/topic/strategy",
		}
		var wg sync.WaitGroup

		wg.Add(len(targets))
		for _, url := range targets {
			go func(u string) {
				defer wg.Done()
				result := checker.CheckURLSync(u)
				if result.Err != nil {
					var unreachable *checker.UnreachableURLError
					if errors.As(result.Err, &unreachable) {
						fmt.Printf("🚫 %s est inaccessible : %v\n", unreachable.URL, unreachable.Err)
					} else {
						fmt.Printf("❌ %s : erreur - %v\n", result.Target, result.Err)
					}
				} else {
					fmt.Printf("✅ %s : OK - %s\n", result.Target, result.Status)
				}
			}(url)
		}

		wg.Wait()
	},
}

// init() est une fonction spéciale de Go, exécutée lors de l'initialisation du package.
func init() {
	// Cette ligne est cruciale : elle "ajoute" la sous-commande `checkCmd` à la commande racine `rootCmd`.
	// C'est ainsi que Cobra sait que 'check' est une commande valide sous 'gowatcher'.
	rootCmd.AddCommand(checkCmd)

	// Ici, vous pouvez ajouter des drapeaux (flags) spécifiques à la commande 'check'.
	// Ces drapeaux ne seront disponibles que lorsque la commande 'check' est utilisée.
	// Exemple (commenté) : checkCmd.Flags().StringVarP(&sourceFile, "source", "s", "", "Fichier contenant les URLs à vérifier")
}
