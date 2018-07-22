package main

import(
	"fmt"
	"sync"
	"time"
)

// 互斥锁基础示例
// 多次解锁会造成painc
/** 
输出:
Lock the lock. (main)
Prepare to lock the lock. (g1)
Prepare to lock the lock. (g2)
Prepare to lock the lock. (g3)
Unlock the lock. (main)
The lock is unlocked. (main)
The lock is locked. (g1)
*/
func main(){
	var mutex sync.Mutex

	fmt.Println("Lock the lock. (main)")
	mutex.Lock()

	for i:=1 ; i<=3; i++ {
		go func(i int){
			fmt.Printf("Prepare to lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (g%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (main)")
	time.Sleep(time.Second)
}