package main
import(
	"fmt"
	"sync"
)

// 互斥锁示例
// 只有走过一个完整的加锁解锁过程才能走下一个
/**
输出:
Lock the lock.
The lock is locked.
Prepare unlock the lock.(1)
The lock is unlocked. (1)
Prepare unlock the lock.(2)
Try to recover the panic.
Recovered the painc("sync: unlock of unlocked mutex").
*/

func main(){
	defer func(){
		fmt.Println("Try to recover the panic.")
		if p:= recover(); p != nil{
			fmt.Printf("Recovered the painc(%#v).\n", p)
		}
	}()

	var mutex sync.Mutex
	fmt.Println("Lock the lock.")
	mutex.Lock()
	fmt.Println("The lock is locked.")

	fmt.Println("Prepare unlock the lock.(1)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (1)")

	fmt.Println("Prepare unlock the lock.(2)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (2)")
	
}
