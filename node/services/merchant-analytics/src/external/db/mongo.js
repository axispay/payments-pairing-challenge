const { MongoClient } = require('mongodb');
async function connectMongo(){const uri=process.env.MONGODB_URI||'mongodb://localhost:27017/payments';const client=new MongoClient(uri);await client.connect();return client.db('payments');}
module.exports={connectMongo};
