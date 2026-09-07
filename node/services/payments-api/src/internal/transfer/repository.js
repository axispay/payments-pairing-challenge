class TransferRepository{constructor(db){this.db=db;this.accounts=db.collection('accounts');this.transactions=db.collection('transactions')}findAccount(id){return this.accounts.findOne({_id:id})}setBalance(id,balance){return this.accounts.updateOne({_id:id},{$set:{balance}})}create(tx){return this.transactions.insertOne(tx)}}
module.exports=TransferRepository;
