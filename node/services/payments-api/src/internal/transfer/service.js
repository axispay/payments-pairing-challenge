class TransferService{constructor(repo,publisher){this.repo=repo;this.publisher=publisher}async transfer(fromAccountId,toAccountId,amount){const from=await this.repo.findAccount(fromAccountId);const to=await this.repo.findAccount(toAccountId);
  if(from.balance<amount){const e=new Error('Insufficient funds');e.status=400;throw e}
  await this.repo.setBalance(fromAccountId,from.balance-amount);await this.repo.setBalance(toAccountId,to.balance+amount);await this.repo.create({fromAccountId,toAccountId,amount,createdAt:new Date()});
  this.publisher.publish('transfer-events',{fromAccountId,toAccountId,amount});return from.balance-amount}}
module.exports=TransferService;
