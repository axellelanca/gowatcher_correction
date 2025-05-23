package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/axellelanca/gowatcher_correction/internal/checker"
)

// ExportResultsToJsonFile prend une liste de checker.ReportEntry et l'écrit dans un fichier JSON.
func ExportResultsToJsonFile(filePath string, results []checker.ReportEntry) error {
	// Utilise json.MarshalIndent pour formater le JSON avec des indentations pour la lisibilité.
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("impossible d'encoder les résultats en JSON: %w", err)
	}

	// Écrit les données JSON dans le fichier spécifié. Les permissions 0644 donnent
	// les droits de lecture/écriture au propriétaire et lecture aux autres.
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire le rapport JSON dans le fichier %s: %w", filePath, err)
	}
	return nil

}
