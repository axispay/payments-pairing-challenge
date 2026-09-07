package com.payments.service;

import com.payments.model.Account;
import com.payments.repository.AccountRepository;
import org.springframework.stereotype.Service;

@Service
public class AccountService {
    private final AccountRepository repository;

    public AccountService(AccountRepository repository) {
        this.repository = repository;
    }

    public Account create(String id, String ownerName, double balance, String merchantId) {
        Account account = new Account();
        account.setId(id);
        account.setOwnerName(ownerName);
        account.setBalance(balance);
        account.setMerchantId(merchantId);
        return repository.create(account);
    }

    public Account get(String id) {
        return repository.find(id);
    }

    public double debit(String id, double amount) throws Exception {
        Account account = repository.find(id);

        if (account == null) {
            throw new IllegalArgumentException("Account not found");
        }

        Thread.sleep(50);

        if (account.getBalance() < amount) {
            throw new InsufficientFundsException();
        }

        double balance = account.getBalance() - amount;
        repository.setBalance(id, balance);

        return balance;
    }

    public static class InsufficientFundsException extends RuntimeException {
    }
}
