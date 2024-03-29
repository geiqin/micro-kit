package pubsub

import (
	"github.com/geiqin/micro-kit/utils"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats/v2"
	"log"
	"os"
)

var myBroker broker.Broker

type EventType string

//消息注册
func init() {
	addr := os.Getenv("MICRO_BROKER_ADDRESS")
	myBroker = nats.NewBroker(broker.Addrs(addr))

	if err := myBroker.Init(); err != nil {
		log.Println("broker init error :", err.Error())
	}

	if err := myBroker.Connect(); err != nil {
		log.Println("broker connect error :", err.Error())
	}

	log.Println("broker register succeed")
}

//消息发布
func Publish(eventName EventType, storeId int64, data string, headers ...map[string]string) error {
	if headers != nil {
		return PublishEvent(string(eventName), storeId, data, headers[0])
	} else {
		return PublishEvent(string(eventName), storeId, data)
	}

}

func PublishEvent(eventName string, storeId int64, data string, headers ...map[string]string) error {
	heads := make(map[string]string)
	if storeId > 0 {
		heads["store_id"] = utils.Int64ToString(storeId)
	}
	if headers != nil {
		for k, v := range headers[0] {
			if k != "store_id" {
				heads[k] = v
			}
		}
	}
	msg := &broker.Message{
		Header: heads,
		Body:   []byte(data),
	}
	err := myBroker.Publish(eventName, msg)
	return err
}

//订阅消息
func SubscribeEvent(eventName string, handler broker.Handler) (broker.Subscriber, error) {
	sub, err := myBroker.Subscribe(eventName, handler)
	return sub, err
}

//订阅消息
func Subscribe(eventName EventType, handler broker.Handler) (broker.Subscriber, error) {
	sub, err := myBroker.Subscribe(string(eventName), handler)
	return sub, err
}
