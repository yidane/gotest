package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"
)

func connectionAndCreateIndex(replace bool) error {
	if err := connection(); err != nil {
		return err
	}
	if err := createIndex(replace); err != nil {
		return err
	}
	fmt.Println("connect es and create index succeed")
	return nil
}

func connection() error {
	var err error
	esClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		return err
	}

	return nil
}

func createIndex(replace bool) error {
	flag, err := esClient.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return err
	}
	if flag && replace {
		deleteIndex, err := esClient.DeleteIndex(indexName).Do(context.TODO())
		if err != nil {
			return err
		}
		if !deleteIndex.Acknowledged {
			return fmt.Errorf("expected ack for creating index; got: %v", deleteIndex.Acknowledged)
		}
		flag = false
	}
	if !flag {
		createIndex, err := esClient.CreateIndex(indexName).BodyJson(indexBody).Do(context.TODO())
		if err != nil {
			return err
		}
		// if createIndex != nil {
		// 	return fmt.Errorf("expected response; got: %v", createIndex)
		// }
		if !createIndex.Acknowledged {
			return fmt.Errorf("expected ack for creating index; got: %v", createIndex.Acknowledged)
		}
	}

	return nil
}
