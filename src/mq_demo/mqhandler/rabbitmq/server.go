package rabbitmq

import(
	"github.com/streadway/amqp"
	"log"
)

// rabbitmq 的相关配置
const(
	RabbitmqHost = "amqp://guest:guest@localhost:5672/"
	QueueName = "hello"
	ConsumerName = ""
	Exchange = "rpc_transaction"
	Durable = false
	DeleteWhenUnused = false
	Exclusive = false
	NoWait = false
	AutoAck = true
	NoLocal = false
	Mandatory = false
	Immediate = false
	DelayExpiration = "5000" // 设置5秒的队列过期时间, 这里仅仅用在延时队列设置当中
)



// Rabbitmq 消息队列
type Rabbitmq struct{}

// Read 向队列读取的方法
func (rb Rabbitmq) Read(f func(jsonStr []byte)){
	conn, err := amqp.Dial(RabbitmqHost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个主监听队列, 延时队列也将会把过期消息转发到这里
	q, err := ch.QueueDeclare(
		QueueName, 			// name
		Durable,   			// durable
		DeleteWhenUnused,   // delete when unused
		Exclusive,   		// exclusive
		NoWait,   			// no-wait
		nil,     			// arguments
	)
	failOnError(err, "Failed to declare a queue")

	declareDelayQueue(ch, QueueName, Exchange)
	declareExchange(ch, Exchange, "fanout")

	// 将主监听队列和 exchange 绑定
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		Exchange, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, 			// queue
		ConsumerName,     	// consumer
		AutoAck,   			// auto-ack
		Exclusive,  		// exclusive
		NoLocal,  			// no-local
		NoWait,  			// no-wait
		nil,    			// args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			go f(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Delay 发送延时消息
func (rb Rabbitmq) Delay(key string,value string, expire string){
	conn, err := amqp.Dial(RabbitmqHost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()


	body := value

	delayName := QueueName + "_delay"
	err = ch.Publish(
		"",     	// exchange
		delayName, 		// routing key
		Mandatory,  	// mandatory
		Immediate,  	// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration: expire,	// 设置五秒的过期时间
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}


// 声明一个延时队列,这个队列不做消费,而是让消息变成死信后再进行转发
func declareDelayQueue(ch *amqp.Channel,channelName string, exchangeName string){
	delayName := channelName + "_delay"
	_, errDelay := ch.QueueDeclare(
		delayName,    	// name
		false, 			// durable
		false, 			// delete when unused
		true,  			// exclusive
		false, 			// no-wait
		amqp.Table{
			"x-dead-letter-exchange":exchangeName,
		},   // arguments
	)
	failOnError(errDelay, "Failed to declare a delay_queue")
}

// 声明一个 exchange, 这里只是为了接收延时队列而设置的一个 exchange
func declareExchange(ch *amqp.Channel, exchangeName string, exType string){
	err := ch.ExchangeDeclare(
		exchangeName,   // name
		exType, // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

// Send 向队列发送的方法
func (rb Rabbitmq) Send(key string,value string){
	conn, err := amqp.Dial(RabbitmqHost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QueueName, 			// name
		Durable,   			// durable
		DeleteWhenUnused,   // delete when unused
		Exclusive,   		// exclusive
		NoWait,   			// no-wait
		nil,     			// arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := value
	err = ch.Publish(
		Exchange,     	// exchange
		q.Name, 		// routing key
		Mandatory,  	// mandatory
		Immediate,  	// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}