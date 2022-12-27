package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	startTime time.Time
	running   bool
	elapsed   time.Duration
}

func (s *Stopwatch) Start() {
	if !s.running {
		s.startTime = time.Now()
		s.running = true
	}
}

func (s *Stopwatch) Stop() {
	if s.running {
		s.elapsed += time.Since(s.startTime)
		s.running = false
	}
}

func (s *Stopwatch) Reset() {
	s.startTime = time.Time{}
	s.running = false
	s.elapsed = 0
}

func (s *Stopwatch) Elapsed() time.Duration {
	if s.running {
		return s.elapsed + time.Since(s.startTime)
	}
	return s.elapsed
}

func main() {
	var stopwatch Stopwatch

	stopwatch.Start()
	time.Sleep(10)
	stopwatch.Stop()
	fmt.Println("Elapsed time:", stopwatch.Elapsed())

	stopwatch.Reset()
	fmt.Println("Elapsed time after reset:", stopwatch.Elapsed())
}
