class AnalyticsRepository{constructor(db){this.collection=db.collection('transactions')}create(transaction){return this.collection.insertOne(transaction)}async find(merchantId,start,end){
  return this.collection.find({merchantId,createdAt:{$gte:start,$lt:end}}).toArray()}}
module.exports=AnalyticsRepository;
