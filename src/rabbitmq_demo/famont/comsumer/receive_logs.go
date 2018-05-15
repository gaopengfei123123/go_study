package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func main() {
	// 建立链接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个主要使用的 exchange
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// 声明一个常规的队列, 其实这个也没必要声明,因为 exchange 会默认绑定一个队列
	q, err := ch.QueueDeclare(
		"test_logs",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 声明一个延时队列, 我们的延时消息就是要发送到这里
	_, errDelay := ch.QueueDeclare(
		"test_delay",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		amqp.Table{
			// 当消息过期时把消息发送到 logs 这个 exchange
			"x-dead-letter-exchange":"logs",
		},   // arguments
	)
	failOnError(errDelay, "Failed to declare a delay_queue")

	err = ch.QueueBind(
		q.Name, // queue name, 这里指的是 test_logs
		"",     // routing key
		"logs", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	// 这里监听的是 test_logs
	msgs, err := ch.Consume(
		q.Name, // queue name, 这里指的是 test_logs
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}