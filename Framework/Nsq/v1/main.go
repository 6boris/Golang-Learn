package v1

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error {
	log.Println("recv:", string(message.Body))
	return nil
}

func testNSQ() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("test", "ch1", nsq.NewConfig())
		if nil != err {
			log.Println(err)
			return
		}

		consumer.AddHandler(&NSQHandler{})
		err = consumer.ConnectToNSQD("10.100.156.207:4150")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	waiter.Wait()
}
