# Analytics

- The goal of this project is to allow us to do everything that segment does(minus the destinations)and to test it locally.  There are several pieces below that will help us do this.

## analytics-platform

- This contains the code to spin up docker containers for kafka(including zookeeper, kafka, schema registry, and kafka connect)

- The goal is to allow us to send data to kafka and have it write out to a file sink(local directory). We want to mimic how kinesis writes out to a firehose ends up in s3.

## analytics-service

- This contains 2 REST endpoints(batchpost and healthcheck) to allow clients to call identify and track.  It uses gorilla mux and kafka-go.

## analytics-test-client

- This contains a test go client to test sending segment events to the analytics service.  This is using the standard analytics-go segment library with a custom url(of the analytics service) to do this.

## analytics-test-website

- This will eventually contain a website that will allow us to test the analytics platform using many different scenarios.

## analytics-plan-runner

- This will contain a client that takes the plan from the test website and executes it.

## Notes

- I know this stack completely misses the event reconciliation part of the workflow so I'm thinking about adding something that will ingest the events into postgresql and then surface some graphs(possibly in cube.js or some other visualization tool).

- In a real world, high event volume stack, I would probably have a process that loads the events into a Snowflake(or other columnar database) and then use a tool like Looker(or tableau) to view the results.

## Version 1 Arch

- Just the basics.

![alt text](./arch_images/AnalyticsPlatformLocal_8_30_19.jpg)

## Version 1 Arch via S3

![alt text](./arch_images/AnalyticsPlatformS3_10_3_19.jpg)

## Final Vision in AWS

- Analytics Platform is black boxed because in theory it could be any event aggregation platform.

![alt text](./arch_images/AnalyticsTestingPlatform_8_23_19.jpg)

