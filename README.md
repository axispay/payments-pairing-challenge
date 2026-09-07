# Backend Payments Challenge

Welcome, and thanks for taking the time to prepare.

This repository contains a small two-service payment system. The interview will be a live
technical session where we work on this codebase together, so we are sharing it a day ahead.

There is nothing to submit and nothing to solve beforehand. The only goal is that you arrive
already familiar with the project: how it is structured, how to run it, and how the main
flows work. That way we can spend the session on the actual coding instead of setup.

## Pick one stack

The same system is implemented three times. Choose whichever language you are most
comfortable with and ignore the other two. They are functionally identical.

| Folder | Stack |
| --- | --- |
| [`go/`](go/) | Go 1.22, Gin, MongoDB driver v2, Sarama / kafka-go |
| [`node/`](node/) | Node 20, Express, MongoDB driver, KafkaJS |
| [`java-spring-boot/`](java-spring-boot/) | Java 21, Spring Boot, MongoTemplate, Spring Kafka |

Each folder has its own `README.md` with the code layout and run instructions for that stack.

## What the system does

Two services, one database, one message broker:

- **payments-api** (port `3001`) owns accounts, debits, payment creation and
  account-to-account transfers. It writes to MongoDB and publishes events to Kafka.
- **merchant-analytics** (port `3002`) stores merchant transactions, consumes the Kafka
  events, and exposes a per-merchant daily totals endpoint.
- **MongoDB** holds accounts, payments, transfers, transactions and analytics events.
- **Kafka** carries `payment-events` and `transfer-events` between the services.

```text
client ──> payments-api ──> MongoDB
               │
               └──> Kafka ──> merchant-analytics ──> MongoDB
                                     ▲
client ──────────────────────────────┘  (REST: transactions, daily totals)
```

### API surface

`payments-api`

| Method | Path | Purpose |
| --- | --- | --- |
| GET | `/health` | liveness |
| POST | `/accounts` | create an account |
| GET | `/accounts/:id` | read an account |
| POST | `/accounts/:id/debit` | debit an account |
| POST | `/payments` | create a payment |
| POST | `/api/transfer` | move money between two accounts |

`merchant-analytics`

| Method | Path | Purpose |
| --- | --- | --- |
| GET | `/health` | liveness |
| POST | `/transactions` | record a merchant transaction |
| GET | `/merchants/:id/daily-totals?date=YYYY-MM-DD` | totals for one merchant on one day |

## How to run

You need Docker with the Compose plugin. Nothing else is required; every stack builds inside
its container.

```bash
cd go            # or: cd node   /   cd java-spring-boot
docker compose up --build
```

Kafka takes a little while to become healthy on first start. The services are ready when both
health checks answer:

```bash
curl http://localhost:3001/health
curl http://localhost:3002/health
```

Then load the sample data (accounts and a day of merchant transactions). The script only uses
the public HTTP endpoints, so it works with any of the three stacks:

```bash
./seed.sh
```

On Windows, run it from Git Bash or WSL, or copy the `curl` calls out of the script.

## How to prepare

An hour or so with the stack you picked is enough. We suggest:

1. **Get it running.** Run `docker compose up --build` and check that both health checks
   answer, so you have seen the services start and know what the running system looks like.
2. **Read the code.** Both services are small. Start from the entry point and follow a request
   through the REST layer, the service layer and the repository layer, and see where Kafka
   comes in.

## During the session

The session is onsite. We will provide a laptop with this project already set up and running,
so you do not need to bring your own. If you prefer to work on your own machine with your own
tools, you are welcome to bring it with the stack you picked already running.

We will pick a part of this codebase and code on it together. Think of it as pair programming
on a real service rather than a quiz. You can look things up and ask questions as you go.

Good luck, and see you soon.
