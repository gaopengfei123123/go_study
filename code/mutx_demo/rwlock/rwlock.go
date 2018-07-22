package main

import(
	"fmt"
	"sync"
	"time"
)

// 读写锁示例
// 写锁只有读锁全解开之后才能进行写锁的操作
// 读锁可以进行并发操作， 而且重复解锁一样会报错
/**
Try to lock for reading... [2]
Locked  for reading... [2]
Try to lock for reading... [1]
Locked  for reading... [1]
Try to lock for reading... [3]
Locked  for reading... [3]
Try to lock for writing....
Try to unlock for reading... [1]
Unlocked  for reading... [1]
Try to unlock for reading... [3]
Unlocked  for reading... [3]
Try to unlock for reading... [2]
Unlocked  for reading... [2]
Unlocked for writing....
*/


func main(){
	var rwm sync.RWMutex
	for i:=1 ; i<=3; i++ {
		go func(i int){
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked  for reading... [%d]\n", i)

			time.Sleep(time.Second * 2)

			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked  for reading... [%d]\n", i)
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try to lock for writing....")
	rwm.Lock()
	fmt.Println("Unlocked for writing....")
}