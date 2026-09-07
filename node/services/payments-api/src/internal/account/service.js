const sleep=(ms)=>new Promise(r=>setTimeout(r,ms));
class AccountService{constructor(repo){this.repo=repo}async create(input){const account={_id:input.id,ownerName:input.ownerName,balance:Number(input.balance),merchantId:input.merchantId||''};await this.repo.create(account);return account}async get(id){return this.repo.findById(id)}async debit(id,amount){const account=await this.repo.findById(id);if(!account){const e=new Error('Account not found');e.status=404;throw e}
  await sleep(50);if(account.balance<amount){const e=new Error('Insufficient funds');e.code='INSUFFICIENT_FUNDS';throw e}const balance=account.balance-amount;await this.repo.setBalance(id,balance);return balance}}
module.exports=AccountService;
