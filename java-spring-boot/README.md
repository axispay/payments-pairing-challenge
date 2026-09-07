# Payments Challenge — Java / Spring Boot

Spring Boot implementation of the two-service payment system described in the
[root README](../README.md).

**Stack:** Java 21 + Spring Boot + MongoTemplate + Spring Kafka + Maven

## Code layout

Both services use a package-first structure:

```text
services/<service>/src/main/java/com/<service>/
├── <Service>Application.java
├── controller/        # REST endpoints, one controller per domain
├── dto/               # request and response objects
├── model/             # domain and MongoDB documents
├── repository/        # all persistence access
├── service/           # all business logic
├── config/            # Spring configuration
└── kafka/             # Kafka producers / consumers
```

For example `payments-api` has `AccountController`, `PaymentController` and
`TransferController`, backed by the matching classes under `service/`.

## Run

```bash
docker compose up --build
```

The first build downloads Maven dependencies and can take a few minutes. Wait until both
services answer:

```bash
curl http://localhost:3001/health
curl http://localhost:3002/health
```

To run a service outside Docker you need JDK 21 and Maven plus MongoDB and Kafka reachable at
`MONGODB_URI` and `KAFKA_BROKERS` (see `docker-compose.yml` for the values used in Compose):

```bash
cd services/payments-api && mvn spring-boot:run
```


## Inspecting the data

MongoDB is exposed on `localhost:27017`, database `payments`:

```bash
docker compose exec mongo mongosh payments --eval 'db.accounts.find().toArray()'
docker compose exec mongo mongosh payments --eval 'db.payments.countDocuments()'
```

Kafka is exposed on `localhost:9092`. To tail a topic:

```bash
docker compose exec kafka /opt/kafka/bin/kafka-console-consumer.sh \
  --bootstrap-server localhost:9092 --topic payment-events --from-beginning
```
