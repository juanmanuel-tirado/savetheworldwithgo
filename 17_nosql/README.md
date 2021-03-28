# 17 NoSQL databases

The examples in this folder assume a standard **Cassandra** service running at 
`localhost:9042` and a the `example` keyspace already available. You can start a basic
cassandra server using a Docker container.

```
docker run -d -p 9042:9042 --name c1 cassandra
docker ps
docker logs c1
```
If everything was OK you should see the logs of your running Cassandra instance.
You need `cqlsh` to be installed in order to create the `example` keyspace.

```
cqlsh -e "create keyspace example with replication = {'class' : 'SimpleStrategy', 'replication_factor':1}"
```

The current examples do not clean the entities created in the example keyspace.
In order to guarantee an example to run, the keyspace must be cleaned.

```
cqlsh -e "drop keyspace example;"
```
