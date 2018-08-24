package main
import(
	"log"
	"net"
	"time"
)

// 模拟 server 出现大量阻塞的场景
// 查看本地支持的最大链接数: sysctl -a|grep kern.ipc.somaxconn
func main(){
	l, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Println("error listen:", err)
		return
	}

	defer l.Close()
	log.Println("listen ok")

	var i int
	for {
		time.Sleep(time.Second * 10)
		if _, err := l.Accept(); err != nil {
			log.Println("accept error:", err)
		}
		i++
		log.Printf("%d: accept a new connection \n", i)
	}

}