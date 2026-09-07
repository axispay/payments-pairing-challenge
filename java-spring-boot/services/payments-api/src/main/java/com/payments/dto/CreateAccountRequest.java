package com.payments.dto;

public record CreateAccountRequest(
        String id,
        String ownerName,
        double balance,
        String merchantId
) {
}
