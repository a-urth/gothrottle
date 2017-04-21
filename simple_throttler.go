package gothrottle

import (
	"sync"
	"time"
)

type SimpleThrottle struct {
	ThrottleParam
	sync.Mutex
	timesMap map[int]int
}

func NewSimpleThrottle(limit int, period time.Duration) *SimpleThrottle {
	return &SimpleThrottle{
		ThrottleParam: ThrottleParam{
			Limit:  limit,
			Period: period,
		},
		timesMap: make(map[int]int),
	}
}

func (st *SimpleThrottle) Record() {
	st.Lock()
	defer st.Unlock()

	var newRecordsNum int
	var currentDelta int
	for {
		current := time.Duration(time.Now().Unix()) * time.Second
		currentDelta = int(current / st.Period)

		// to keep memory footprint shallow delete unrelevant records since
		// their time already passed
		for k := range st.timesMap {
			if k < currentDelta {
				delete(st.timesMap, k)
			}
		}

		recordsNum, ok := st.timesMap[currentDelta]
		if !ok {
			newRecordsNum = 1
			break
		}

		if recordsNum <= st.Limit {
			newRecordsNum = recordsNum + 1
			break
		}

		// Actually there is no reason to sleep whole period, we need to sleep
		// till next time frame, but since time duration mangling in go not that
		// simple and for sake of simplicity its like it is
		time.Sleep(st.Period)
	}

	st.timesMap[currentDelta] = newRecordsNum
}
