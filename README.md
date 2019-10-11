# dirtytest
For run kafka cluster use command `docker-compose up -d` from directory that contains docker-compose.yml file

Before run producer/consumer programs create topic with this command `docker run --net=host --rm confluentinc/cp-kafka:latest kafka-topics --create --topic topic2 --partitions 3 --replication-factor 3 --if-not-exists --zookeeper localhost:12181  --config min.insync.replicas=2`

For run producer use command `./main` from directory producer/bin/

For run consumer that will be read messages from last offset pass any argument to program. For example, `./main start`
For run consumer that will be read messages from first message in topic pass keyword `first` as argument to program. For example, `./main first`

