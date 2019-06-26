package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"strconv"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			decodeConfig := nsq.NewConfig()
			c, err := nsq.NewConsumer("test1", "test-channel1", decodeConfig)
			if err != nil {
				log.Panic("Could not create consumer")
			}
			c.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
				log.Println(strconv.Itoa(i)+" received:", string(msg.Body))
				return nil
			}))

			err = c.ConnectToNSQD("nsq.kyle.link:4150")
			if err != nil {
				log.Panicln("connect error:", err.Error())
			}
			select {}
		}(i)
	}

	wg.Wait()
}
