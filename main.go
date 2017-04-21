package gothrottle

import "time"

type Throttler interface {
	Record()
	GetParams() ThrottleParam
}

type ThrottleParam struct {
	Limit  int
	Period time.Duration
}

func (tp ThrottleParam) GetParams() ThrottleParam {
	return tp
}
