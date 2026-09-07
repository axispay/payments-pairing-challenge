package com.payments.repository;

import com.payments.model.Payment;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class PaymentRepository {
    private final MongoTemplate mongo;

    public PaymentRepository(MongoTemplate mongo) {
        this.mongo = mongo;
    }

    public Payment create(Payment payment) {
        return mongo.insert(payment, "payments");
    }
}
