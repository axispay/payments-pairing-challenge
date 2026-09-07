package com.analytics.kafka;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import java.util.Date;
import java.util.Map;

@Component
public class PaymentEventConsumer {
    private final MongoTemplate mongo;
    private final ObjectMapper mapper = new ObjectMapper();

    public PaymentEventConsumer(MongoTemplate mongo) {
        this.mongo = mongo;
    }

    @KafkaListener(topics = {"payment-events", "transfer-events"}, groupId = "merchant-analytics")
    public void consume(ConsumerRecord<String, String> record) {
        Object payload;

        try {
            payload = mapper.readValue(record.value(), Map.class);
        } catch (Exception e) {
            payload = Map.of("raw", record.value());
        }

        mongo.insert(
                Map.of(
                        "topic", record.topic(),
                        "payload", payload,
                        "receivedAt", new Date()
                ),
                "analytics_events"
        );
    }
}
