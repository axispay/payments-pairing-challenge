const { Kafka } = require('kafkajs');
class KafkaProducer{constructor(){this.producer=new Kafka({clientId:'payments-api',brokers:(process.env.KAFKA_BROKERS||'localhost:9092').split(',')}).producer()}async connect(){await this.producer.connect()}async publish(topic,payload){await this.producer.send({topic,messages:[{value:JSON.stringify(payload)}]})}}
module.exports={KafkaProducer};
