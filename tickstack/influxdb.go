package tickstack

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

func Query(url, measurement, ifql string) {
	config := client.HTTPConfig{
		Addr: url,
	}
	c, err := client.NewHTTPClient(config)
	if err != nil {
		log.Fatal("connect error: ", err)
	}
	defer c.Close()

	q := client.Query{
		Command:  ifql,
		Database: measurement,
	}

	if response, err := c.Query(q); err == nil && response.Error() == nil {
		for _, result := range response.Results {
			for _, message := range result.Messages {
				log.Printf("level: %s; text: %s", message.Level, message.Text)
			}

			for _, row := range result.Series {
				log.Println("Name: ", row.Name)
				log.Println("Partial: ", row.Partial)
				log.Println("Tags: ", row.Tags)
				for _, col := range row.Columns {
					log.Println(col)
				}
				for _, value := range row.Values {
					log.Println(value)
				}
			}
		}
	} else {
		log.Fatal(response.Error())
	}
}
