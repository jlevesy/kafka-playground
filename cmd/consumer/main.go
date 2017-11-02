package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/shopify/sarama"
)

var (
	brokers = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "Kafka brokers to connect to as a comma separated list")
)

func main() {
	flag.Parse()

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(*brokers, ",")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer producer.Close()

	for i := 0; i < 100; i++ {
		producer.SendMessage(
			&sarama.ProducerMessage{
				Topic: "interesting",
			},
		)
	}

}
