package checker

import (
	"fmt"
	"net/http"
	"time"
)

// CheckResult Représente le résultat d'une vérification
type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckURL(url string, results chan<- CheckResult) {
	// Timeout court pour éviter de bloquer trop longtemps
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		results <- CheckResult{
			Target: url,
			Err:    fmt.Errorf("Request failed:  %w", err),
		}
		return
	}
	defer resp.Body.Close()

	results <- CheckResult{
		Target: url,
		Status: resp.Status,
	}
}
