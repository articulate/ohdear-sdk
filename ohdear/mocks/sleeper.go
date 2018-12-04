package mocks

import "time"

type MockSleeper struct {
	SleepCall struct {
		Receives struct {
			Time float64
		}
	}
}

func (s MockSleeper) Sleep(seconds time.Duration) {
	s.SleepCall.Receives.Time = seconds.Seconds()
}
