package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		log.Println(req.Body)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
			Value: body,
		}
		err = kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func main() {
	// get kafka writer using environment variables.
	kafkaURL := "localhost:32886"
	topic := "events"
	kafkaWriter := getKafkaWriter(kafkaURL, topic)

	defer kafkaWriter.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/v1/batch", batchpost(kafkaWriter)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
