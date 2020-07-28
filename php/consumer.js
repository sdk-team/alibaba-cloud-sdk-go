
const Kafka = require('node-rdkafka');
const config = require('./setting');
console.log(Kafka.features);
console.log(Kafka.librdkafkaVersion);
console.log(config)

let kafkaConsumer = new Kafka.KafkaConsumer({
    /*'debug': 'all',*/
    'api.version.request': 'true',
    'bootstrap.servers': config['bootstrap_servers'],
    'message.max.bytes': 32000,
    'fetch.message.max.bytes': 32000,
    'max.partition.fetch.bytes': 32000,
    'security.protocol': 'sasl_ssl',
    'ssl.ca.location': config['ca'],
    'sasl.mechanisms': 'PLAIN',
    'sasl.username': config['sasl_plain_username'],
    'sasl.password': config['sasl_plain_password'],
    'group.id': config['consumer_id']
  }, {
    'auto.offset.reset': config['reset'],
  });
  

// Flowing mode:
kafkaConsumer.connect();

kafkaConsumer.on('ready', function () {
    console.log("connect ok");

    kafkaConsumer.subscribe([config['php_topic_name']]);

    // Consume from the librdtesting-01 topic. This is what determines
    // the mode we are running in. By not specifying a callback (or specifying
    // only a callback) we get messages as soon as they are available.
    kafkaConsumer.consume();
})

kafkaConsumer.on('data', function (data) {
    // Output the actual message contents
    console.log(data);
});


kafkaConsumer.on('event.log', function (event) {
    console.log("event.log", event);
});

kafkaConsumer.on('error', function (error) {
    console.log("error:" + error);
});

kafkaConsumer.on('event', function (event) {
    console.log("event:" + event);
});


exports.getTask = async function () {
    return await new Promise((resolve, reject) => {
        function clearUp() {
            kafkaConsumer.removeListener('data', onData);
            kafkaConsumer.removeListener('error', onError);
        }

        function onData(data) {
            // Output the actual message contents
            let task = data.value;
            console.log(`${(new Date).toString()} get new task: ${task}`);
            clearUp();
            resolve(JSON.parse(task));
        }

        function onError(error) {
            console.log("error:" + error);
            clearUp();
            reject(error);
        }
        kafkaConsumer.on('data', onData);

        kafkaConsumer.on('error', onError);

    })
}