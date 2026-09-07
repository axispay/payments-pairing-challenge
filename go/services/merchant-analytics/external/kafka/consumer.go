package kafka

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Consumer struct {
	brokers []string
	db      *mongo.Database
}

func NewConsumer(raw string, db *mongo.Database) *Consumer {
	brokers := []string{"localhost:9092"}
	if raw != "" {
		brokers = strings.Split(raw, ",")
	}

	return &Consumer{
		brokers: brokers,
		db:      db,
	}
}

func (c *Consumer) Start() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup(c.brokers, "merchant-analytics", config)
	if err != nil {
		log.Error().Err(err).Msg("failed to create Kafka consumer group")
		return
	}
	defer group.Close()

	handler := &consumerGroupHandler{db: c.db}
	topics := []string{"payment-events", "transfer-events"}

	for {
		if err := group.Consume(context.Background(), topics, handler); err != nil {
			log.Error().Err(err).Msg("Kafka consume failed")
			time.Sleep(time.Second)
		}
	}
}

type consumerGroupHandler struct {
	db *mongo.Database
}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var payload any
		if err := json.Unmarshal(message.Value, &payload); err != nil {
			payload = map[string]any{"raw": string(message.Value)}
		}

		_, err := h.db.Collection("analytics_events").InsertOne(session.Context(), map[string]any{
			"topic":      message.Topic,
			"payload":    payload,
			"receivedAt": time.Now(),
		})
		if err != nil {
			log.Error().Err(err).Str("topic", message.Topic).Msg("failed to store analytics event")
			continue
		}

		session.MarkMessage(message, "")
	}

	return nil
}
