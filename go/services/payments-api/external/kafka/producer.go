package kafka

import (
	"context"
	"strings"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct{ brokers []string }

func NewProducer(raw string) *Producer {
	brokers := []string{"localhost:9092"}
	if raw != "" {
		brokers = strings.Split(raw, ",")
	}
	return &Producer{brokers: brokers}
}

func (p *Producer) Publish(ctx context.Context, topic string, value []byte) error {
	writer := &kafkago.Writer{Addr: kafkago.TCP(p.brokers...), Topic: topic, Balancer: &kafkago.LeastBytes{}, WriteTimeout: 5 * time.Second}
	defer writer.Close()
	return writer.WriteMessages(ctx, kafkago.Message{Value: value})
}
func (p *Producer) Close() error { return nil }
