package singal

import "sync"

type SingalM struct {
	Name  string
	State string
}

type Singal struct {
	//sc chan *SingalM
	wg *sync.WaitGroup
}

func NewSingal(len int) *Singal {
	/*
		if len == 0 {
			return &Singal{sc: make(chan *SingalM)}
		} else {
			return &Singal{sc: make(chan *SingalM, len)}
		}
	*/
	return &Singal{wg: &sync.WaitGroup{}}
}

func (s *Singal) singal(name string, state string) {
	/*
		s.sc <- &SingalM{
			Name:  name,
			State: state,
		}
	*/

	//s.wg.Done()
}

func (s *Singal) Done() {
	/*
		s.sc <- &SingalM{
			Name:  name,
			State: state,
		}
	*/

	s.wg.Done()
}

func (s *Singal) Add() {
	s.wg.Add(1)
}

func (s *Singal) Wait() {
	s.wg.Wait()
}
