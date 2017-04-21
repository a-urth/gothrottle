package gothrottle

import (
	"testing"
	"time"
)

func TestSimpleThrottleSleep(t *testing.T) {
	limit := 3
	period := time.Duration(5) * time.Second
	throttlers := []Throttler{
		NewChannelThrottler(limit, period),
		NewSimpleThrottle(limit, period),
	}
	for _, throttler := range throttlers {
		now := time.Now()
		for i := 0; i < (throttler.GetParams().Limit*2)+1; i++ {
			throttler.Record()
		}

		passed := time.Now().Sub(now)
		if passed < throttler.GetParams().Period {
			t.Errorf("Less time passed then period - %v", passed)
		}
	}
}

func TestSimpleThrottleNoSleep(t *testing.T) {
	limit := 10
	period := time.Duration(5) * time.Second
	throttlers := []Throttler{
		NewChannelThrottler(limit, period),
		NewSimpleThrottle(limit, period),
	}
	for _, throttler := range throttlers {
		now := time.Now()
		for i := 0; i < throttler.GetParams().Limit; i++ {
			throttler.Record()
		}

		time.Sleep(throttler.GetParams().Period)

		for i := 0; i < throttler.GetParams().Limit; i++ {
			throttler.Record()
		}

		passed := time.Now().Sub(now)
		if passed > throttler.GetParams().Period+time.Duration(1)*time.Second {
			t.Errorf("More time passed then expected - %v", passed)
		}
	}
}
