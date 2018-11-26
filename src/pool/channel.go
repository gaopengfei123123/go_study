package pool

import (
	// "fmt"
	"errors"
	"sync"
	"time"
)

//Config 初始化配置
type Config struct {
	MaxConn   int
	AliveConn int
	Factory   func() (interface{}, error)
	Ping      func(interface{}) error
	Close     func(interface{}) error
	TimeOut   time.Duration
}

type channelPool struct {
	mux     sync.Mutex
	conns   chan *connChan
	factory func() (interface{}, error)
	close   func(interface{}) error
	ping    func(interface{}) error
	timeOut time.Duration
}

type connChan struct {
	conn interface{}
	t    time.Time
}

// NewPool 初始化连接池
func NewPool(config Config) (pool Pool, err error) {
	if config.MaxConn <= 0 || config.AliveConn < 0 || config.AliveConn > config.MaxConn {
		err = errors.New("invalid conn config")
		return
	}

	if config.Factory == nil {
		err = errors.New("we need factory func")
		return
	}

	cp := &channelPool{
		conns:   make(chan *connChan, config.MaxConn),
		factory: config.Factory,
		close:   config.Close,
		timeOut: config.TimeOut,
		ping:    config.Ping,
	}

	for i := 0; i < config.AliveConn; i++ {
		conn, err := cp.factory()
		if err != nil {
			err = errors.New("factory make conn error")
			break
		}
		cp.conns <- &connChan{conn: conn, t: time.Now()}
	}

	return cp, err
}

func (c *channelPool) Get() (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	select{
	case c.conns ->:
		
	}

	return nil, nil
}

func (c *channelPool) Put(interface{}) error {
	return nil
}
func (c *channelPool) Close(interface{}) error {
	return nil
}
func (c *channelPool) Ping(interface{}) error {

	return nil
}
func (c *channelPool) Len() int {
	return cap(c.conns)
}

func (c *channelPool) Release() {

}
