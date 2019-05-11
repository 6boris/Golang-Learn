package v1

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
	"testing"
)

func TestNSQHandler_HandleMessage(t *testing.T) {
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
		err = consumer.ConnectToNSQD("127.0.0.1:4150")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("test", "ch1", nsq.NewConfig())
		if nil != err {
			log.Println(err)
			return
		}

		consumer.AddHandler(&NSQHandler{})
		err = consumer.ConnectToNSQD("120.79.134.74:4150")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	waiter.Wait()
}
