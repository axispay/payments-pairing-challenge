package com.analytics.repository;

import com.analytics.model.TransactionRecord;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Repository;

import java.util.Date;
import java.util.List;

@Repository
public class AnalyticsRepository {
    private final MongoTemplate mongo;

    public AnalyticsRepository(MongoTemplate mongo) {
        this.mongo = mongo;
    }

    public TransactionRecord create(TransactionRecord transaction) {
        return mongo.insert(transaction, "transactions");
    }

    public List<TransactionRecord> find(String merchantId, Date start, Date end) {
        return mongo.find(
                Query.query(
                        Criteria.where("merchantId")
                                .is(merchantId)
                                .and("createdAt")
                                .gte(start)
                                .lt(end)
                ),
                TransactionRecord.class,
                "transactions"
        );
    }
}
