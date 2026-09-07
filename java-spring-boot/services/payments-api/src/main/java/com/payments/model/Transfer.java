package com.payments.model;

import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;

@Document("transactions")
public record Transfer(
        String fromAccountId,
        String toAccountId,
        double amount,
        Date createdAt
) {
}
