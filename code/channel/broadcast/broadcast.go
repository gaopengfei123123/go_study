package main

import "context"

import "fmt"

import "log"

import "sync"

// BroadService chan广播
type BroadService struct {
	regist    chan interface{}
	broadList []chan interface{}
	remove    chan interface{}
	BroadChan chan interface{}
	Ctx       context.Context
	Cancel    context.CancelFunc
}

// NewBroadcast 初始化广播队列
func NewBroadcast() (service *BroadService) {
	ctx, cancel := context.WithCancel(context.Background())
	service = &BroadService{
		regist:    make(chan interface{}),
		BroadChan: make(chan interface{}, 1),
		remove:    make(chan interface{}),
		broadList: make([]chan interface{}, 2),
		Ctx:       ctx,
		Cancel:    cancel,
	}
	return
}

// Listen 监听广播
func (th *BroadService) Listen() (ch chan interface{}, cch <-chan struct{}) {
	ch = make(chan interface{})
	cch = th.Ctx.Done()
	th.regist <- ch
	return
}

// Speak 发送广播
func (th *BroadService) Speak(msg interface{}) {
	th.BroadChan <- msg
}

// Remove 移除
func (th *BroadService) Remove(ch chan interface{}) error {
	th.remove <- ch
	return nil
}

// Run 运行
func (th *BroadService) Run() {
	for {
		select {
		case message := <-th.BroadChan:
			for i := 0; i < len(th.broadList); i++ {
				if th.broadList[i] != nil {
					th.broadList[i] <- message
				}
			}
		case ch := <-th.regist:
			th.broadList = append(th.broadList, ch.(chan interface{}))
		case ch := <-th.remove:
			for i := 0; i < len(th.broadList); i++ {
				if th.broadList[i] == ch.(chan interface{}) || th.broadList[i] == nil {
					fmt.Printf("移除: %#+v \n", ch)
					th.broadList = append(th.broadList[:i], th.broadList[i+1:]...)
					continue
				}
			}
		}
	}
}

func main() {
	log.Println("start")
	server := NewBroadcast()

	go server.Run()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	listen1, cancel1 := server.Listen()
	go test(listen1, cancel1, "线程1", wg)

	wg.Add(1)
	listen2, cancel2 := server.Listen()
	go test(listen2, cancel2, "线程2", wg)

	for i := 0; i < 10; i++ {
		server.Speak(i)
	}
	server.Cancel()

	wg.Wait()

	fmt.Println("done")

}

func test(ch <-chan interface{}, cancel <-chan struct{}, sign string, wg *sync.WaitGroup) {
	fmt.Printf("channel: %s \n", sign)
LOOP:
	for {
		select {
		case message := <-ch:
			fmt.Printf("chan: %s, message: %#+v \n", sign, message)
		case <-cancel:
			fmt.Printf("chan: %s, break \n", sign)
			break LOOP
		}
	}
	wg.Done()
}
