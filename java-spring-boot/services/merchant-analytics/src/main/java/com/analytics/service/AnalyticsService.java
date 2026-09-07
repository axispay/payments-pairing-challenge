package com.analytics.service;

import com.analytics.dto.DailyTotalsResponse;
import com.analytics.model.TransactionRecord;
import com.analytics.repository.AnalyticsRepository;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class AnalyticsService {
    private final AnalyticsRepository repository;

    public AnalyticsService(AnalyticsRepository repository) {
        this.repository = repository;
    }

    public TransactionRecord createTransaction(String merchantId, double amount, String status, Date createdAt) {
        TransactionRecord transaction = new TransactionRecord();
        transaction.setMerchantId(merchantId);
        transaction.setAmount(amount);
        transaction.setStatus(status);
        transaction.setCreatedAt(createdAt == null ? new Date() : createdAt);
        return repository.create(transaction);
    }

    public DailyTotalsResponse daily(String id, Date start, Date end) {
        List<TransactionRecord> transactions = repository.find(id, start, end);

        double total = 0;
        int count = 0;
        Map<String, Integer> byStatus = new HashMap<>();

        for (TransactionRecord transaction : transactions) {
            total += transaction.getAmount();
            count++;

            if (transaction.getStatus() != null) {
                byStatus.merge(transaction.getStatus(), 1, Integer::sum);
            }
        }

        return new DailyTotalsResponse(total, count, byStatus);
    }
}
