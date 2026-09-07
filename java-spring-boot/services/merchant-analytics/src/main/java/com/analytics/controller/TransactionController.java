package com.analytics.controller;

import com.analytics.dto.CreateTransactionRequest;
import com.analytics.model.TransactionRecord;
import com.analytics.service.AnalyticsService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TransactionController {
    private final AnalyticsService service;

    public TransactionController(AnalyticsService service) {
        this.service = service;
    }

    @PostMapping("/transactions")
    public ResponseEntity<TransactionRecord> create(@RequestBody CreateTransactionRequest request) {
        TransactionRecord transaction = service.createTransaction(
                request.merchantId(),
                request.amount(),
                request.status(),
                request.createdAt()
        );
        return ResponseEntity.status(201).body(transaction);
    }
}
