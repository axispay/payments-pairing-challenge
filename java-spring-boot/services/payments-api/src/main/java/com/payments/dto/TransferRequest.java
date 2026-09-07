package com.payments.dto;

public record TransferRequest(
        String fromAccountId,
        String toAccountId,
        double amount
) {
}
