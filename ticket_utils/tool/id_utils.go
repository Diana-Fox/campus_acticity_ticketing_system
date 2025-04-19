package tool

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch       = int64(1609459200000) // 设置一个起始时间戳，这里使用2021-01-01 00:00:00 UTC
	workerBits  = 10
	numberBits  = 12
	workerMax   = -1 ^ (-1 << workerBits)
	numberMax   = -1 ^ (-1 << numberBits)
	timeShift   = workerBits + numberBits
	workerShift = numberBits
	timeMask    = -1 ^ (-1 << timeShift)
	workerMask  = -1 ^ (-1 << workerBits)
	numberMask  = -1 ^ (-1 << numberBits)
)

var (
	workerId      int64
	number        int64
	lastTimestamp int64 = epoch // 声明并初始化 lastTimestamp
	mutex         sync.Mutex
)

// InitWorker 初始化工作节点ID
func InitWorker(id int64) error {
	if id < 0 || id > workerMax {
		return errors.New("worker ID must be between 0 and " + string(workerMax))
	}
	workerId = id
	number = 0
	lastTimestamp = epoch // 初始化 lastTimestamp
	return nil
}

// NextId 生成下一个唯一的ID
func NextId() (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().UnixNano() / 1000000 // 获取当前时间戳（毫秒）
	if now < lastTimestamp {
		return 0, errors.New("invalid system clock")
	}

	if now == lastTimestamp {
		number = (number + 1) & numberMask
		if number == 0 {
			for now <= lastTimestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		number = 0
	}

	lastTimestamp = now

	// 将时间戳、workerId 和序列号拼接成一个64位的ID
	return (now-epoch)<<timeShift | (workerId << workerShift) | number, nil
}