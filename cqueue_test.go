package cqueue

import (
	"log"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

	var a = 1
	log.Printf("%p", a)
	return
	Run()

	AddQueue("a", 10, false)
	AddQueue("b", 10, true)
	AddQueue("c", 10, false)

	AddTask("a", func() {
		log.Println("a 开始")
		time.Sleep(3 * time.Second)
		log.Println("a 结束")
	})

	AddTask("a", func() {
		log.Println("a 开始")
		time.Sleep(3 * time.Second)
		log.Println("a 结束")
	})

	time.Sleep(1 * time.Second)
	DelQueue("a")

	AddTask("a", func() {
		log.Println("a 开始")
		time.Sleep(3 * time.Second)
		log.Println("a 结束")
	})

	AddTask("b", func() {
		log.Println("b 开始")
		time.Sleep(5 * time.Second)
		log.Println("b 结束")
	})
	AddTask("b", func() {
		log.Println("b 开始")
		time.Sleep(5 * time.Second)
		log.Println("b 结束")
	})
	AddTask("b", func() {
		log.Println("b 开始")
		time.Sleep(5 * time.Second)
		log.Println("b 结束")
	})
	AddTask("b", func() {
		log.Println("b 开始")
		time.Sleep(5 * time.Second)
		log.Println("b 结束")
	})
	AddTask("c", func() {
		log.Println("c 开始")
		time.Sleep(2 * time.Second)
		log.Println("c 结束")
	})
	AddTask("c", func() {
		log.Println("c 开始")
		time.Sleep(2 * time.Second)
		log.Println("c 结束")
	})

	log.Println("over")
	time.Sleep(10 * time.Second)
}
