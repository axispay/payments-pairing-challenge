package com.payments.dto;

public record CreatePaymentRequest(
        String accountId,
        double amount,
        String currency
) {
}
