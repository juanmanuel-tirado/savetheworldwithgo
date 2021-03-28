#Kafka

These examples assume Kafka instance is running and it is  available at 
`localhost:9092`. You can follow the 
the [official guide](https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html#ce-docker-quickstart) 
to have a basic Kafka deployment in containers.

A minimal Docker compose file is provided here. Notice that only the services
required by the examples are deployed in the provided Docker compose file. To
start the environment run:

```
docker-compose up -d
```

This will start the containers and check all the dependencies. All the containers
must be up and ready to have a working environment. In order to check that 
everything is ready you can run:
```
docker ps
CONTAINER ID   IMAGE                                   COMMAND                  CREATED        STATUS       PORTS                                            NAMES
32155ffabe37   confluentinc/cp-kafka-rest:6.1.0        "/etc/confluent/dock…"   18 hours ago   Up 4 hours   0.0.0.0:8082->8082/tcp                           rest-proxy
3bb8cfafd9ee   confluentinc/cp-schema-registry:6.1.0   "/etc/confluent/dock…"   18 hours ago   Up 4 hours   0.0.0.0:8081->8081/tcp                           schema-registry
344a9836ad27   confluentinc/cp-server:6.1.0            "/etc/confluent/dock…"   18 hours ago   Up 4 hours   0.0.0.0:9092->9092/tcp, 0.0.0.0:9101->9101/tcp   broker
5458ec512452   confluentinc/cp-zookeeper:6.1.0         "/etc/confluent/dock…"   18 hours ago   Up 4 hours   2888/tcp, 0.0.0.0:2181->2181/tcp, 3888/tcp       zookeeper
```

If you have any trouble with this procedure, refer to the official documentation.
