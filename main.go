package main

import (
	"context"
	"github-pr-status/internal"
	"github.com/cli/cli/v2/pkg/iostreams"
)

func main() {
	options := internal.ParseOptions()
	queries := []internal.Query{
		{"Created", "is:pr is:open author:%s sort:updated-desc"},
		{"Requested(not draft)", "is:pr is:open archived:false -is:draft review-requested:%s sort:updated-desc"},
	}

	ctx := context.Background()
	cs := iostreams.System().ColorScheme()

	results, prCounts := internal.FetchPrData(ctx, options.Token, options.Users, queries)
	if options.Verbose {
		internal.PrintPrData(options.Users, queries, results, cs)
	}
	internal.PrintPrSummary(options.Users, queries, prCounts, cs)
}
