class PaymentService{constructor(repo,publisher){this.repo=repo;this.publisher=publisher}async create(input){const payment={...input,status:'charged',createdAt:new Date()};
  await this.repo.create(payment);await this.publisher.publish('payment-events',{type:'PAYMENT_CREATED',payment});return payment}}
module.exports=PaymentService;
