package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error,msg string){
	if err != nil{
		log.Fatalf("%s:%s",msg,err)
	}
}

func main(){
	// 连接到rabbitmq服务器
	conn,err := amqp.Dial("amqp://lsx_01:admin123@127.0.0.1:5672/my_vhost")
	failOnError(err,"Failed connect!")
	defer conn.Close()

	// 配置连接套接字，用于定义连接协议，以及身份验证
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列用以向其发送消息
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 推送消息
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}