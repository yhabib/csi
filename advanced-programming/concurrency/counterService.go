package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counterService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently without any additional synchronization.
	getNext() uint64
}

type naiveService struct {
	counter uint64
}

// sometimes will work sometimes it will not
func (s *naiveService) getNext() uint64 {
	s.counter++
	return s.counter
}

type atomicService struct {
	counter uint64
}

func (s *atomicService) getNext() uint64 {
	return atomic.AddUint64(&s.counter, 1)
}

type mutexService struct {
	counter uint64
	mu      sync.Mutex
}

func (s *mutexService) getNext() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	return s.counter
}

// Launch a separate goroutine with exclusive access to a private counter value;
// handle getNext() calls by making “requests” and receiving “responses” on two separate channels

type goService struct {
	req chan<- struct{}
	res <-chan uint64
}

func newGoService() *goService {
	req := make(chan struct{})
	res := make(chan uint64)
	counter := uint64(0)
	go func(c uint64) {
		for {
			<-req
			c++
			res <- c
		}
	}(counter)
	return &goService{req, res}
}

func (s *goService) getNext() uint64 {
	s.req <- struct{}{}
	return <-s.res
}

func main() {
	numOfGoRoutines := 50
	numOfIterationsPerGoRoutine := 10
	var maxValue uint64 = uint64(numOfGoRoutines * numOfIterationsPerGoRoutine)

	services := make(map[string]counterService)
	services["naive"] = &naiveService{}
	services["atomic"] = &atomicService{}
	services["mutex"] = &mutexService{}
	services["channel"] = newGoService()

	for k := range services {
		// From GoPL 8.5 -> We wait for each service to finish
		var wg sync.WaitGroup

		for i := 0; i < numOfGoRoutines; i++ {
			wg.Add(1)
			go func(k string) {
				defer wg.Done()
				// Shouldn't this increase monotonically
				var prevValue uint64 = 0
				for j := 0; j < numOfIterationsPerGoRoutine; j++ {
					newValue := services[k].getNext()
					if newValue-prevValue != 1 {
						fmt.Print("No monotonically incremented")
					}
					prevValue = newValue
				}
			}(k)
		}

		wg.Wait()
		maxVal := services[k].getNext() - 1
		if maxVal != maxValue {
			fmt.Printf("%s service FAILED %d != %d\n", k, maxVal, maxValue)
		} else {
			fmt.Printf("%s service DONE %d == %d\n", k, maxVal, maxValue)
		}
	}
}
