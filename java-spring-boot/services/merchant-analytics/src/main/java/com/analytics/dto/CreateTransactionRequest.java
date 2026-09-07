package com.analytics.dto;

import java.util.Date;

public record CreateTransactionRequest(
        String merchantId,
        double amount,
        String status,
        Date createdAt
) {
}
