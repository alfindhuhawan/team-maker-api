package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"team-maker-api/shared/infrastructure/config"
	"team-maker-api/shared/model/errorenum"
	"team-maker-api/shared/model/payload"

	amqp "github.com/rabbitmq/amqp091-go"
)

//const exchangeName = "simple.exchange"
//const exchangeType = amqp.ExchangeTopic

var exchangeName = "delayed.exchange"
var exchangeType = "x-delayed-message"

type publisherImpl struct {
	rabbitMQChannel *amqp.Channel
	cfg             *config.Config
}

func getURL(cfg *config.Config) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.MessageBroker.Rabbitmq.Username,
		cfg.MessageBroker.Rabbitmq.Password,
		cfg.MessageBroker.Rabbitmq.Address,
		cfg.MessageBroker.Rabbitmq.Port,
	)
}

// NewPublisher is
// url
func NewPublisher(cfg *config.Config) Publisher {

	url := getURL(cfg)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer ch.Close()
	fmt.Println(ch)

	return &publisherImpl{
		rabbitMQChannel: nil,
		cfg:             cfg,
	}
}

// Publish is
func (m *publisherImpl) Publish(topic string, delayInMS int, data payload.Payload) error {

	url := getURL(m.cfg)

	conn, errConn := amqp.Dial(url)
	if errConn != nil {
		fmt.Println(errConn)
		// log.Fatal(errConn.Error())
		return errorenum.RabbitMQError
	}
	defer conn.Close()

	if errConn == nil {
		rabbitMQChannel, errChan := conn.Channel()
		if errChan != nil {
			fmt.Println(errChan.Error())
			// log.Fatal(errChan.Error())
		}
		defer rabbitMQChannel.Close()

		if errChan == nil {
			fmt.Println(rabbitMQChannel)
			dataInBytes, err := json.Marshal(data)
			if err != nil {
				return err
			}

			headers := amqp.Table{
				"x-delay": delayInMS, // only for x-delay-message
			}

			err = rabbitMQChannel.Publish(
				exchangeName, // exchange
				topic,        // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        dataInBytes,
					Headers:     headers,
				})
			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}

type subscriberImpl struct {
	url       string
	queueName string
	topicMap  map[string]HandleFunc
}

// NewSubscriber is
func NewSubscriber(queueName string, cfg *config.Config) Subscriber {
	url := getURL(cfg)
	return &subscriberImpl{
		url:       url,
		queueName: queueName,
		topicMap:  map[string]HandleFunc{},
	}
}

func (r *subscriberImpl) Handle(topic string, onReceived HandleFunc) {

	r.topicMap[topic] = onReceived

}

// Run is
// "amqp://guest:guest@localhost:5672/"
func (r *subscriberImpl) Run() {

	conn, err := amqp.Dial(r.url)
	if err != nil {
		panic(err.Error())
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	rabbitMQChannel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		err := rabbitMQChannel.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	args := amqp.Table{
		"x-delayed-type": "topic", // only for x-delay-message
	}

	err = rabbitMQChannel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		args,         // arguments
	)
	if err != nil {
		panic(err.Error())
	}

	for s := range r.topicMap {

		q, err := rabbitMQChannel.QueueDeclare(
			r.queueName+"-"+s, // name
			false,             // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // no-wait
			nil,               // arguments
		)
		if err != nil {
			panic(err.Error())
		}

		err = rabbitMQChannel.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			exchangeName, // exchange
			false,
			nil,
		)
		if err != nil {
			panic(err.Error())
		}

		deliveryMsg, err := rabbitMQChannel.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("[RabbitMQ] Queue:%-50s Event:%s\n", q.Name, s)

		go func(routingKey string) {
			for d := range deliveryMsg {
				var data payload.Payload
				err := json.Unmarshal(d.Body, &data)
				r.topicMap[routingKey](data, err)
				//log.Printf("recv %s %s", d.RoutingKey, data.Data)
			}
		}(s)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

}
