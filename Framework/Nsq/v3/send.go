package v3

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

var producer *nsq.Producer

func StartSend(n int) {
	//ip := "127.0.0.1:4150"
	ip := "nsq.kyle.link:4152"
	InitProducer(ip)
	defer producer.Stop()

	for i := 0; i < n; i++ {
		//time.Sleep(time.Microsecond * 100000)
		err := Publish("test", RandString(100000))
		if err == nil {
			log.Println("NSQ success send: ", i)
		} else {
			log.Println("NSQ fail send: ", i)
		}
	}

}

func InitProducer(str string) {
	var err error
	fmt.Println("addrss: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())

	if err != nil {
		panic("producer is nil")
	}
}

//发布消息
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return nil
}
