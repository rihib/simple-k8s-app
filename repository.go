package main

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type Result struct {
	Word  string
	Count int
}

func getAPICounts() int {
	ctx := context.Background()

	projectID := "tron-151603"
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	query := client.Query(`
        SELECT word, COUNT(*) as count
        FROM ` + "`bigquery-public-data.samples.shakespeare`" + `
        GROUP BY word
        ORDER BY count DESC
        LIMIT 1;
    `)
	it, err := query.Read(ctx)
	if err != nil {
		log.Fatalf("query.Read: %v", err)
	}

	var result Result
	if err := it.Next(&result); err == iterator.Done {
		return 0
	} else if err != nil {
		log.Fatalf("iterator.Next: %v", err)
	}

	return result.Count
}
