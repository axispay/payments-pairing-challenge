package com.payments.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.payments.model.Payment;
import com.payments.kafka.EventPublisher;
import com.payments.repository.PaymentRepository;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.Map;

@Service
public class PaymentService {
    private final PaymentRepository repository;
    private final EventPublisher publisher;
    private final ObjectMapper mapper = new ObjectMapper();

    public PaymentService(PaymentRepository repository, EventPublisher publisher) {
        this.repository = repository;
        this.publisher = publisher;
    }

    public Payment create(String accountId, double amount, String currency) throws Exception {
        Payment payment = new Payment();
        payment.setAccountId(accountId);
        payment.setAmount(amount);
        payment.setCurrency(currency);
        payment.setStatus("charged");
        payment.setCreatedAt(new Date());

        payment = repository.create(payment);

        publisher.publishAndWait(
                "payment-events",
                mapper.writeValueAsString(Map.of("type", "PAYMENT_CREATED", "payment", payment))
        );

        return payment;
    }
}
