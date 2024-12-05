package utils

import "time"

type Clock struct {
	startTime time.Time
	endTime   time.Time
	handler   func(delta time.Duration)
}

func NewClock(handler func(delta time.Duration)) *Clock {
	currentTime := time.Now()
	return &Clock{
		startTime: currentTime,
		endTime:   currentTime,
		handler:   handler,
	}
}

func (clock *Clock) Delta() time.Duration {
	return clock.endTime.Sub(clock.startTime)
}

func (clock *Clock) Stop() {
	clock.endTime = time.Now()
	clock.handler(clock.Delta())
}
