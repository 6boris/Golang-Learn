package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {

	//config := nsq.NewConfig()
	//p, err := nsq.NewProducer("nsq.kyle.link:4250", config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		//发布一条消息
	//		for i := 0; i < 100000000; i++ {
	//			err = p.Publish("test1", []byte(time.Now().String()))
	//			if err != nil {
	//				log.Panicln(err.Error())
	//			}
	//		}
	//	}()
	//}


	initProducer("nsq.kyle.link:4150","test2")
	initProducer("nsq.kyle.link:4250","test2")
	initProducer("nsq.kyle.link:4350","test2")

	select {}


}

func initProducer(addr string, topic string) {

	config := nsq.NewConfig()
	p, err := nsq.NewProducer(addr, config)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		go func() {
			//发布一条消息
			for i := 0; i < 100000000; i++ {
				err = p.Publish(topic, []byte(time.Now().String()))
				if err != nil {
					log.Panicln(err.Error())
				}
			}
		}()
	}

}
