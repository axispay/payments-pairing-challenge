package com.payments.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.payments.model.Account;
import com.payments.model.Transfer;
import com.payments.kafka.EventPublisher;
import com.payments.repository.TransferRepository;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.Map;

@Service
public class TransferService {
    private final TransferRepository repository;
    private final EventPublisher publisher;
    private final ObjectMapper mapper = new ObjectMapper();

    public TransferService(TransferRepository repository, EventPublisher publisher) {
        this.repository = repository;
        this.publisher = publisher;
    }

    public double transfer(String fromId, String toId, double amount) throws Exception {
        Account from = repository.find(fromId);
        Account to = repository.find(toId);

        if (from.getBalance() < amount) {
            throw new IllegalArgumentException("Insufficient funds");
        }

        repository.setBalance(fromId, from.getBalance() - amount);
        repository.setBalance(toId, to.getBalance() + amount);
        repository.create(new Transfer(fromId, toId, amount, new Date()));

        String event = mapper.writeValueAsString(Map.of(
                "fromAccountId", fromId,
                "toAccountId", toId,
                "amount", amount
        ));

        publisher.publish("transfer-events", event);

        return from.getBalance() - amount;
    }
}
