package main
import(
	"log"
	"net"
)

func main(){
	l , err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Println("listen error: " , err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error: " , err)
			return
		}
		log.Println("accept a new connection")
		go handler(c)
	}
}

func handler(c net.Conn){
	defer c.Close()

	for {
		var buf = make([]byte, 10)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}