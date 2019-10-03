# api-service

- This contains 2 REST endpoints(batchpost and healthcheck) to allow clients to call identify and track.  It uses gorilla mux and kafka-go.
- This is an extremely basic api that posts events to kafka.
- Build it using: docker build -t api-service .
- Run it using docker-compose in the parent directory.
