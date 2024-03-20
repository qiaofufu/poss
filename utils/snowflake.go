package utils

import (
	"strconv"
	"sync"
	"time"
)

const (
	startTime     = 1768433720000
	numberMask    = (1 << 12) - 1
	workIDMask    = ((1 << 22) - 1) ^ ((1 << 12) - 1)
	timestampMask = ((1 << 63) - 1) ^ ((1 << 22) - 1)
)

type SnowFlake struct {
	workID    int64
	number    int64
	timestamp int64
	mu        sync.Mutex
}

func NewSnowFlake(workID int64) *SnowFlake {
	if workID > 1<<10-1 {
		panic("workID too large")
	}
	return &SnowFlake{workID: workID}
}

func (s *SnowFlake) GenerateID() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	timestamp := time.Now().UnixMilli()
	if timestamp == s.timestamp {
		s.number++
	} else {
		s.number = 0
	}
	s.timestamp = timestamp

	id := (((s.timestamp - startTime) << 22) & timestampMask) | (((s.workID) << 12) & workIDMask) | (s.number & numberMask)
	return strconv.FormatInt(id, 10)
}
