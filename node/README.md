# Payments Challenge — Node.js

Node.js implementation of the two-service payment system described in the
[root README](../README.md).

**Stack:** Node 20 + Express + MongoDB driver + KafkaJS

## Code layout

Both services share the same layout:

```text
services/<service>/src/
├── index.js           # entry point: wiring, HTTP server, Kafka
├── api/rest/          # app.js, routes.js, handlers/
├── internal/          # business domains: model + repository + service
│   ├── account/       # payments-api only
│   ├── payment/       # payments-api only
│   ├── transfer/      # payments-api only
│   └── analytics/     # merchant-analytics only
└── external/          # MongoDB and Kafka infrastructure
    ├── db/
    └── kafka/
```

## Run

```bash
docker compose up --build
```

Wait until both services answer:

```bash
curl http://localhost:3001/health
curl http://localhost:3002/health
```

To run a service outside Docker you need Node 20 plus MongoDB and Kafka reachable at
`MONGODB_URI` and `KAFKA_BROKERS` (see `docker-compose.yml` for the values used in Compose):

```bash
cd services/payments-api && npm install && npm start
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
