package main

import (
	"github.com/jfeng45/gmessaging"
	"github.com/jfeng45/gmessaging/cmd"
	"log"
	"runtime"
)
type PaymentEvent struct {
	Id int
	SourceAccount string
}

func main() {
	testEncoded()
	//testNonEncoded()
}
func testEncoded() {
	subject :="testEncoded"
	mi, err := cmd.InitMessagingEncodedService()
	//defer mi.Close()
	//defer mi.CloseConnection()
	if err != nil {
		log.Println("err:", err)
	}
	_, err = mi.Subscribe(subject, func(pe PaymentEvent) {
		log.Println("payload:",pe)
		mi.Close()
		mi.CloseConnection()
	})
	if err != nil {
		log.Println("err:",err)
	}
	log.Printf("Listening on [%s]", subject)
	runtime.Goexit()
}

func testNonEncoded() {
	subject :="testNonEncoded"
	mi, err := cmd.InitMessagingService()
	//defer mi.Close()
	if err != nil {
		log.Println("err:", err)
	}
	_, err = mi.Subscribe(subject, func(msg *gmessaging.GMessage) {
		if msg != nil {
			log.Println("payload:",string(msg.Data))
		} else {
			log.Println("payload is nil")
		}
		mi.Close()
	})
	if err != nil {
		log.Println("err:",err)
	}
	log.Printf("Listening on [%s]", subject)
	runtime.Goexit()
}


