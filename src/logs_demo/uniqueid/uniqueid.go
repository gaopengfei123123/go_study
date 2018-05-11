package uniqueid

import(
	"fmt"
	"sync"
	"errors"
	"time"
)

func Test(){
	fmt.Println("uniqueid")
}

const (
	workerBits uint8 = 10
	numberBits uint8 = 12
	workerMax int64 = -1 ^ (-1 << workerBits)
	numberMax int64 = -1 ^ (-1 << numberBits)
	epoch int64 = 1525104000000 // 以 2018/5/1 零点为时间零点, 续一口命, 防止出现2038年的悲剧

	workerShift uint8 = workerBits + numberBits
	timeShift   uint8 = numberBits

)


// Worker 生成 id 的基本参数
type Worker struct{
	mu sync.Mutex
	timestamp int64
	workerID  int64
	number    int64
}

// NewWorker 创建一个实例
func NewWorker(wid int64) (*Worker, error){
	if wid < 0 || wid > workerMax{
		return nil, errors.New("Worker ID excess of quantify")
	}

	return &Worker{
		timestamp:0,
		workerID: wid,
		number: 0,
	}, nil
}

// GetID 生成 uniqueID
func (w *Worker) GetID() int64{
	// 加上互斥锁,防止并发冲突
	w.mu.Lock()
	defer w.mu.Unlock()

	// 获取毫秒数
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++

		if w.number > numberMax {
			// 这里用 for 而不是 sleep 是为了获得一个完整的毫秒
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	// 利用位偏移和或操作将参数合并, 从左向右依次为:0 位为正负值恒为0 1~42 为时间戳, 43~52 为机器 id, 53~64 为 number 数
	ID := int64( ((now-epoch) << timeShift) | (w.workerID << workerShift) |  (w.number))
	return ID
}