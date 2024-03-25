package main

import (
	"context"
	"log"
	"os"

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

	queryBytes, err := os.ReadFile("query.sql")
	if err != nil {
		log.Fatalf("ioutil.ReadFile: %v", err)
	}
	query := client.Query(string(queryBytes))
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
