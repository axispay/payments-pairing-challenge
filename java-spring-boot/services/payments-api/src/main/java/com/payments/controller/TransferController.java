package com.payments.controller;

import com.payments.dto.TransferRequest;
import com.payments.service.TransferService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
public class TransferController {
    private final TransferService service;

    public TransferController(TransferService service) {
        this.service = service;
    }

    @PostMapping("/api/transfer")
    public ResponseEntity<?> transfer(@RequestBody TransferRequest request) throws Exception {
        double balance = service.transfer(
                request.fromAccountId(),
                request.toAccountId(),
                request.amount()
        );

        return ResponseEntity.ok(Map.of("success", true, "newBalance", balance));
    }
}
