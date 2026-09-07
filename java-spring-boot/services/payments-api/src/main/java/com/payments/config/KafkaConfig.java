package com.payments.config;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
public class KafkaConfig {
    @Bean
    public NewTopic paymentEventsTopic() {
        return TopicBuilder.name("payment-events").build();
    }

    @Bean
    public NewTopic transferEventsTopic() {
        return TopicBuilder.name("transfer-events").build();
    }
}
