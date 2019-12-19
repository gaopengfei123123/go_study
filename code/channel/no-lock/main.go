package main

import (
	"log"
	"time"
)

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {

	for {
		w, ok := <-in
		log.Printf("%v: %v \n", w, ok)
		if !ok {
			close(out)
			break
		}
		w.z = w.x + w.y
		time.Sleep(time.Second * 1)
		out <- w

	}
}

func Run() {
	in, out := make(chan *Work), make(chan *Work)

	workerNum := 4
	for i := 0; i < workerNum; i++ {
		go worker(in, out)
	}
	go sendWorks(in)
	receveWorks(out)
}

func sendWorks(in chan<- *Work) {
	log.Println("开始输入")
	for i := 0; i < 10; i++ {
		in <- &Work{
			x: i,
			y: 1,
			z: 0,
		}
	}
	log.Println("输入完毕")
	close(in)
}

func receveWorks(out <-chan *Work) {
	log.Println("开始输出")
	for w := range out {
		log.Println("输出:", w.z)
	}
	log.Println("输出完毕")
}

func main() {
	Run()
}
