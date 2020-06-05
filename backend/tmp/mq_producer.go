package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	"log"
)



func createService() micro.Service {
	service := micro.NewService(
		micro.Broker(kafka.NewBroker(func(o *broker.Options) {
			o.Addrs = []string{"34.92.78.87:9092"}
		})),
	)
	if err := broker.Connect(); err != nil {
		log.Fatal(err)
	}
	service.Init()
	return service
}

//func main() {
//	srv := createService()
//	router := httprouter.New()
//	srv.Server().Handle(srv.Server().NewHandler(router))
//	broker.Publish("Topic主题", &broker.Message{
//		Header: map[string]string{
//			"AAA": "BBBBB",
//			"CCCCC": "DDDDDD",
//		},
//		Body: []byte("消息内容"),
//	})
//	broker.Subscribe("Topic主题", func(p broker.Publication) error {
//		brokerHeader := p.Message().Header
//		aaa := brokerHeader["AAA"]
//		bbb := string(p.Message().Body)
//	})
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//}