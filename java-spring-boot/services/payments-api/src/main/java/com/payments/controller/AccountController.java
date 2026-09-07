package com.payments.controller;

import com.payments.dto.CreateAccountRequest;
import com.payments.dto.DebitAccountRequest;
import com.payments.model.Account;
import com.payments.service.AccountService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
public class AccountController {
    private final AccountService service;

    public AccountController(AccountService service) {
        this.service = service;
    }

    @PostMapping("/accounts")
    public ResponseEntity<Account> create(@RequestBody CreateAccountRequest request) {
        Account account = service.create(
                request.id(),
                request.ownerName(),
                request.balance(),
                request.merchantId()
        );
        return ResponseEntity.status(201).body(account);
    }

    @GetMapping("/accounts/{id}")
    public ResponseEntity<?> get(@PathVariable String id) {
        Account account = service.get(id);

        if (account == null) {
            return ResponseEntity.status(404).body(Map.of("error", "Account not found"));
        }

        return ResponseEntity.ok(account);
    }

    @PostMapping("/accounts/{id}/debit")
    public ResponseEntity<?> debit(
            @PathVariable String id,
            @RequestBody DebitAccountRequest request
    ) throws Exception {
        try {
            return ResponseEntity.ok(Map.of("balance", service.debit(id, request.amount())));
        } catch (AccountService.InsufficientFundsException e) {
            return ResponseEntity.status(500).body(Map.of("error", "Insufficient funds"));
        }
    }
}
