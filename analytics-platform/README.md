# analytics-platform

## How to start things up

1. Go to the analytics-service directory first and build that service.
- This is the go web service that pushes data to kafka.

2. Run: docker-compose up -d
- This starts up the stack(zookeeper, kafka, kafka connect, schema registry, and analytics service)

3. docker exec analytics-platform_kafka-connect_1 mkdir -p /tmp/quickstart/file
- This creates the mountpoint so we can see the output of the file sink

4. docker exec analytics-platform_kafka-connect_1 touch /tmp/quickstart/file/output.txt
- This creates the actual file the output of the file sink will go into

5. curl -X POST -H "Content-Type: application/json" --data '{"name": "quickstart-file-sink", "config": {"connector.class":"FileStreamSink", "tasks.max":"1", "topics":"events", "file": "/tmp/quickstart/file/output.txt", "name": "quickstart-file-sink"}}' http://localhost:8083/connectors
- This creates the file sink in kafka connect.  In practice this would be part of the container build but here we have the rest api exposed so we can tinker.

6. curl -s -X GET http://localhost:8083/connectors/quickstart-file-sink/status
- This is used to view the status of the sink.

7. Run: docker-compose down
- This destroys the stack(zookeeper, kafka, kafka connect, schema registry, and analytics service)