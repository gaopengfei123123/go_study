package pool

// Pool 连接池对外暴露的接口
type Pool interface {
	Get() (interface{}, error)
	Put(interface{}) error
	Close(interface{}) error
	Release()
	Len() int
	Ping(interface{}) error
}
