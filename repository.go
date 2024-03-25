package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type Result struct {
	Count int
}

func getAPICounts() ([]Result, error) {
	ctx := context.Background()

	projectID := "tron-151603"
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	queryBytes, err := os.ReadFile("query.sql")
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", err)
	}

	query := client.Query(string(queryBytes))
	it, err := query.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("query.Read: %v", err)
	}

	var results []Result
	for {
		var result Result
		err := it.Next(&result)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("iterator.Next: %v", err)
		}
		results = append(results, result)
	}

	return results, nil
}
