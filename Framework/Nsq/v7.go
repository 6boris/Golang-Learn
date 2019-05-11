package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

const (
	ModeAll = iota
	ModeRoundRobin
)

type Publisher interface {
	Publish(string, []byte) error
}

type PublishHandler struct {
	Publisher
	addresses util.StringArray
	counter   uint64
	mode      int
}

func (ph *PublishHandler) HandleMessage(m *nsq.Message) error {
	switch ph.mode {
	case ModeAll:
		for _, addr := range ph.addresses {
			err := ph.Publish(addr, m.Body)
			if err != nil {
				return err
			}
		}
	case ModeRoundRobin:
		idx := ph.counter % uint64(len(ph.addresses))
		err := ph.Publish(ph.addresses[idx], m.Body)
		if err != nil {
			return err
		}
		ph.counter++
	}

	return nil
}

type PostPublisher struct{}

func (p *PostPublisher) Publish(addr string, msg []byte) error {
	buf := bytes.NewBuffer(msg)
	resp, err := HttpPost(addr, buf)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("got status code %d", resp.StatusCode))
	}
	return nil
}

func main() {
	r, err := nsq.NewReader(*topic, *channel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for i := 0; i < *numPublishers; i++ {
		handler := &PublishHandler{
			Publisher: &PostPublisher{},
			addresses: postAddrs,
			mode:      mode,
		}
		r.AddHandler(handler)
	}

	for _, addrString := range lookupdHTTPAddrs {
		log.Printf("lookupd addr %s", addrString)
		err := r.ConnectToLookupd(addrString)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	<-r.ExitChan
}
