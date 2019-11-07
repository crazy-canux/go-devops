package tickstack

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

// Query from influxdb.
func Query(url, measurement, ifql string) ([]string, [][]interface{}, error) {
	config := client.HTTPConfig{
		Addr: url,
	}
	c, err := client.NewHTTPClient(config)
	if err != nil {
		log.Println("Connect failed.")
		return nil, nil, err
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
				//for _, col := range row.Columns {
				//	log.Println(col)
				//}
				//for _, value := range row.Values {
				//	log.Println(value)
				//}
				return row.Columns, row.Values, nil
			}
		}
	} else {
		log.Println("Query failed.")
		return nil, nil, err
	}
	return nil, nil, nil
}
