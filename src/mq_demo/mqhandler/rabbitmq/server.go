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
	Exchange = ""
	Durable = false
	DeleteWhenUnused = false
	Exclusive = false
	NoWait = false
	AutoAck = true
	NoLocal = false
	Mandatory = false
	Immediate = false
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

	q, err := ch.QueueDeclare(
		QueueName, 			// name
		Durable,   			// durable
		DeleteWhenUnused,   // delete when unused
		Exclusive,   		// exclusive
		NoWait,   			// no-wait
		nil,     			// arguments
	)
	failOnError(err, "Failed to declare a queue")

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
func (rb Rabbitmq) Delay(key string,value string){

}

// Send 向队列发送的方法
func (rb Rabbitmq) Send(key string,value string){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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