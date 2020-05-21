package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	kafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// healthCheck -- To provide a response for the healthcheck
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alive")
}

// batchpost -- To handle identify calls from the client
func batchpost(producer *kafka.Producer, topic string) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}

		deliveryChan := make(chan kafka.Event)

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          body,
			Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		}, deliveryChan)
		if err != nil {
			log.Print(err)
		}

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			log.Print(m.TopicPartition.Error)
		} else {
			log.Printf("info", "Delivered message", "topic", *m.TopicPartition.Topic, "partition", m.TopicPartition.Partition, "offset", m.TopicPartition.Offset)
		}
	})
}

// get access to kafka
func getKafkaProducer(kafkaURL string) *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaURL})
	if err != nil {
		log.Fatalln(err)
	}

	return producer
}

func main() {
	kafkaURL := os.Getenv("KAFKA_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	log.Println("Running with the following parameters:")
	log.Printf("KAFKA_URL: %s", kafkaURL)
	log.Printf("KAFKA_TOPIC: %s", kafkaTopic)

	// get kafka writer using environment variables.
	kafkaProducer := getKafkaProducer(kafkaURL)

	defer kafkaProducer.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/v1/batch", batchpost(kafkaProducer, kafkaTopic)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
