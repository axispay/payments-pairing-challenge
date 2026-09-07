package com.payments.kafka;

import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class EventPublisher {
    private final KafkaTemplate<String, String> kafka;

    public EventPublisher(KafkaTemplate<String, String> kafka) {
        this.kafka = kafka;
    }

    public void publish(String topic, String payload) {
        kafka.send(topic, payload);
    }

    public void publishAndWait(String topic, String payload) throws Exception {
        kafka.send(topic, payload).get();
    }
}
