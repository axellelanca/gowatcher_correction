package checker

import (
	"net/http"
	"time"
)

// CheckResult Représente le résultat d'une vérification
type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckURLSync(url string) CheckResult {
	// Timeout court pour éviter de bloquer trop longtemps
	client := http.Client{Timeout: 3 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			Target: url,
			Err:    &UnreachableURLError{URL: url, Err: err},
		}
	}
	defer resp.Body.Close()

	return CheckResult{Target: url, Status: resp.Status}
}
