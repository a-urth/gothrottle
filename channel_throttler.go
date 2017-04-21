package gothrottle

import (
	"log"
	"time"
)

type ChannelThrottler struct {
	ThrottleParam
	c chan bool
}

func NewChannelThrottler(limit int, period time.Duration) *ChannelThrottler {
	throttler := ChannelThrottler{
		ThrottleParam: ThrottleParam{
			Limit:  limit,
			Period: period,
		},
		c: make(chan bool, limit),
	}

	duration := time.Duration(period.Nanoseconds() / int64(limit))
	go func(records chan bool, ticker *time.Ticker) {
		for _ = range ticker.C {
			<-records
		}
	}(throttler.c, time.NewTicker(duration))

	return &throttler
}

func (ct *ChannelThrottler) Record() {
	ct.c <- true
	log.Println(len(ct.c))
}
