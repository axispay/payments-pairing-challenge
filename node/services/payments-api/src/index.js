const { connectMongo } = require('./external/db/mongo');
const { KafkaProducer } = require('./external/kafka/producer');
const { createApp } = require('./api/rest/app');
const { createRouter } = require('./api/rest/routes');
const AccountRepository = require('./internal/account/repository');
const AccountService = require('./internal/account/service');
const PaymentRepository = require('./internal/payment/repository');
const PaymentService = require('./internal/payment/service');
const TransferRepository = require('./internal/transfer/repository');
const TransferService = require('./internal/transfer/service');

async function main() {
  const db = await connectMongo();
  const publisher = new KafkaProducer(); await publisher.connect();
  const accountService = new AccountService(new AccountRepository(db));
  const paymentService = new PaymentService(new PaymentRepository(db), publisher);
  const transferService = new TransferService(new TransferRepository(db), publisher);
  const app = createApp(createRouter({ accountService, paymentService, transferService }));
  app.listen(process.env.PORT || 3001, () => console.log('payments-api listening'));
}
main().catch((err) => { console.error(err); process.exit(1); });
