package mocks

import (
	"time"
)

type MockSleeper struct {
	SleepCall struct {
		Receives struct {
			Time float64
		}
		Count int
	}
}

func (s *MockSleeper) Sleep(seconds time.Duration) {
	s.SleepCall.Count = s.SleepCall.Count + 1
}
