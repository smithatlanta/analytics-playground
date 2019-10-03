package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	kafka "github.com/segmentio/kafka-go"
)

// healthCheck -- To provide a response for the healthcheck
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alive")
}

// batchpost -- To handle identify calls from the client
func batchpost(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
			Value: body,
		}
		log.Println(msg)

		err = kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}

// get access to kafka
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func main() {
	kafkaURL := os.Getenv("KAFKA_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")

	log.Println("Running with the following parameters:")
	log.Printf("KAFKA_URL: %s", kafkaURL)
	log.Printf("KAFKA_TOPIC: %s", kafkaTopic)

	// get kafka writer using environment variables.
	kafkaWriter := getKafkaWriter(kafkaURL, kafkaTopic)

	defer kafkaWriter.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/v1/batch", batchpost(kafkaWriter)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
