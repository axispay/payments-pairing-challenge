package com.payments.controller;

import com.payments.dto.CreatePaymentRequest;
import com.payments.model.Payment;
import com.payments.service.PaymentService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PaymentController {
    private final PaymentService service;

    public PaymentController(PaymentService service) {
        this.service = service;
    }

    @PostMapping("/payments")
    public ResponseEntity<Payment> create(@RequestBody CreatePaymentRequest request) throws Exception {
        Payment payment = service.create(request.accountId(), request.amount(), request.currency());
        return ResponseEntity.status(201).body(payment);
    }
}
