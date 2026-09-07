class AccountRepository{constructor(db){this.collection=db.collection('accounts')}create(account){return this.collection.insertOne(account)}findById(id){return this.collection.findOne({_id:id})}async setBalance(id,balance){await this.collection.updateOne({_id:id},{$set:{balance}})}}
module.exports=AccountRepository;
