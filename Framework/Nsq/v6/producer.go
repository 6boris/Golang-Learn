package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("nsq.kyle.link:4152", config)
	if err != nil {
		panic(err)
	}
	//发布一条消息
	for i := 0; i < 100000000; i++ {
		err = p.Publish("test", []byte(time.Now().String()))
		if err != nil {
			log.Panicln(err.Error())
		}
	}

}
