package state_machine

import (
	"sync"
	"time"
)

const (
	StatePending = "pending" //未开始
	StateRunning = "running" //执行中
	StateFailed  = "failed"  //执行失败
	StateSuccess = "success" //执行成功
)

type state struct {
	start  time.Time
	end    time.Time
	err    error
	state  string
	rwlock sync.RWMutex
}

// NewState creates a new State with the given name and initial State.
func newState() *state {
	return &state{
		state:  StatePending,
		rwlock: sync.RWMutex{},
		err:    nil,
	}
}

// GetState returns the current State.
func (s *state) GetState() string {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.state
}

func (s *state) GetStart() time.Time {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.start
}

func (s *state) GetEnd() time.Time {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.end
}

// SetState sets the State to the new State.
func (s *state) Success() {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	// Check if the current State is not already StateSuccess or StateFailed
	if s.state == StatePending || s.state == StateRunning {
		s.state = StateSuccess
		s.end = time.Now()
	}
}

func (s *state) Running() {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	//取消、运行中、失败，都不可以被设置，但是可以重复执行一个成功节点
	if s.state == StatePending || s.state == StateSuccess {
		s.state = StateRunning
		s.start = time.Now()
	}
}

func (s *state) Failed(err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	// Check if the current State is not already StateSuccess or StateFailed
	s.state = StateFailed
	s.end = time.Now()
	s.err = err
}

// IsPending checks if the State is pending.
func (s *state) IsPending() bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.state == StatePending
}

// IsRunning checks if the State is running.
func (s *state) IsRunning() bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.state == StateRunning
}

// IsFailed checks if the State is failed.
func (s *state) IsFailed() bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.state == StateFailed
}

func (s *state) Error() error {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.err
}

func (s *state) IsSuccess() bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return s.state == StateSuccess
}
