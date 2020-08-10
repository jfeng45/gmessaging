package main

import (
	"github.com/jfeng45/gmessaging/cmd"
	"log"
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
	p := PaymentEvent{1, "sourceAccount"}
	mi, err := cmd.InitMessagingEncodedService()
	if err != nil {
		log.Println("err:", err)
	}
	err = mi.Publish(subject, p)
	if err != nil {
		log.Println("err:", err)
	}
	err = mi.Flush()
	if err != nil {
		log.Println("err:", err)
	}
	mi.Close()
	mi.CloseConnection()
}
func testNonEncoded() {
	subject :="testNonEncoded"
	//subject :="test"
	msg := "123"
	msgArr := []byte(msg)
	mi, err := cmd.InitMessagingService()
	if err != nil {
		log.Println("err:", err)
	}
	err = mi.Publish(subject, msgArr)
	if err != nil {
		log.Println("err:", err)
	}
	err = mi.Flush()
	if err != nil {
		log.Println("err:", err)
	}
	mi.Close()
}


