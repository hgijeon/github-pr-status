package internal

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/tableprinter"
)

func PrintPrData(users []string, queries []Query, results map[string]map[string]SearchResult, cs *iostreams.ColorScheme) {
	for _, user := range users {
		userResults := results[user]
		limit := 10
		fmt.Printf("\n\n### %s ###\n", user)
		for qDescription, result := range userResults {
			fmt.Printf("\n%s (TOP %d)\n", qDescription, limit)

			tp := newTablePrinterWithHeaders(os.Stdout, []string{"Updated At", "Title", "User"}, cs)
			for i, pr := range result.Items {
				if i >= limit {
					break
				}
				tp.AddField(pr.UpdatedAt.Format(time.DateTime))
				tp.AddField(pr.Title, tableprinter.WithColor(func(s string) string { return createLink(s, pr.URL) }))
				tp.AddField(pr.User.Login)
				tp.EndRow()
			}

			if err := tp.Render(); err != nil {
				log.Fatalf("Error rendering table for query '%s': %v", qDescription, err)
			}
		}
	}
}

func PrintPrSummary(users []string, queries []Query, prCounts map[string]map[string]int, cs *iostreams.ColorScheme) {
	fmt.Println("\n\n### PR Counts Summary ###")
	tp := newTablePrinterWithHeaders(os.Stdout, append([]string{"User"}, func() []string {
		headers := make([]string, len(queries))
		for i, q := range queries {
			headers[i] = q.Description
		}
		return headers
	}()...), cs)

	for _, user := range users {
		counts := prCounts[user]
		tp.AddField(user)
		for _, q := range queries {
			searchURL := fmt.Sprintf("https://github.com/pulls?q=%s", url.QueryEscape(fmt.Sprintf(q.Query, user)))
			tp.AddField(fmt.Sprintf("%d", counts[q.Description]), tableprinter.WithColor(func(s string) string { return createLink(s, searchURL) }))
		}
		tp.EndRow()
	}

	if err := tp.Render(); err != nil {
		log.Fatalf("Error rendering PR counts summary table: %v", err)
	}
}
