package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {

	config := nsq.NewConfig()
	p, err := nsq.NewProducer("nsq.kyle.link:4152", config)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 10; i++ {
		err = p.Publish("My_NSQ_Topic", []byte("sample NSQ message:"+strconv.Itoa(i)))
		if err != nil {
			log.Panic(err)
		}
	}
}

func Producer() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("nsq.kyle.link:4152", config)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		err = p.Publish("My_NSQ_Topic", []byte("sample NSQ message:"+strconv.Itoa(i)))
		if err != nil {
			log.Panic(err)
		}
	}
}
