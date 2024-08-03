package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type PullRequest struct {
	Title string `json:"title"`
	User  struct {
		Login string `json:"login"`
	} `json:"user"`
	URL       string    `json:"html_url"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SearchResult struct {
	TotalCount int           `json:"total_count"`
	Items      []PullRequest `json:"items"`
}

type Query struct {
	Description string
	Query       string
}

func FetchPrData(ctx context.Context, token string, users []string, queries []Query) (map[string]map[string]SearchResult, map[string]map[string]int) {
	prCounts := make(map[string]map[string]int)
	results := make(map[string]map[string]SearchResult)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, user := range users {
		prCounts[user] = make(map[string]int)
		results[user] = make(map[string]SearchResult)
		for _, q := range queries {
			wg.Add(1)
			go func(user string, q Query) {
				defer wg.Done()
				query := fmt.Sprintf(q.Query, user)
				result, err := fetchPullRequests(ctx, token, query)
				if err != nil {
					log.Fatalf("Error fetching pull requests for query '%s': %v", q.Description, err)
				}
				mu.Lock()
				prCounts[user][q.Description] = result.TotalCount
				results[user][q.Description] = result
				mu.Unlock()
			}(user, q)
		}
	}

	wg.Wait()
	return results, prCounts
}

func fetchPullRequests(ctx context.Context, token, query string) (SearchResult, error) {
	apiURL := fmt.Sprintf("https://api.github.com/search/issues?q=%s", url.QueryEscape(query))
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return SearchResult{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SearchResult{}, fmt.Errorf("error executing request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return SearchResult{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return SearchResult{}, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
