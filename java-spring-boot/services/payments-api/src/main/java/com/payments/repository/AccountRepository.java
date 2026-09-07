package com.payments.repository;

import com.payments.model.Account;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.stereotype.Repository;

@Repository
public class AccountRepository {
    private final MongoTemplate mongo;

    public AccountRepository(MongoTemplate mongo) {
        this.mongo = mongo;
    }

    public Account create(Account account) {
        return mongo.insert(account, "accounts");
    }

    public Account find(String id) {
        return mongo.findById(id, Account.class, "accounts");
    }

    public void setBalance(String id, double balance) {
        mongo.updateFirst(
                Query.query(Criteria.where("_id").is(id)),
                Update.update("balance", balance),
                "accounts"
        );
    }
}
