package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"crawler/engine"
	"github.com/pkg/errors"
	"crawer_distributed/config"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item "+"#%d: %v", itemCount, item)
			itemCount++

			err := Save(client, item)
			if err != nil {
				log.Printf("Item Saver: error "+"saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(config.ElasticIndex).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
