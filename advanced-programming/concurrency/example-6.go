package main

import (
	"fmt"
	"sync"
)

type coordinator struct {
	lock   sync.RWMutex
	leader string
}

func newCoordinator(leader string) *coordinator {
	return &coordinator{
		lock:   sync.RWMutex{},
		leader: leader,
	}
}

func (c *coordinator) logState() {
	c.lock.RLock()
	defer c.lock.RUnlock()

	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) logUnsafeState() {
	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.leader = leader

	if shouldLog {
		c.logUnsafeState()
	}
}

func main() {
	c := newCoordinator("us-east")
	c.logState()
	c.setLeader("us-west", true)
}

// Problem: Deadlock -> setLeader acquires the lock but logState tries to lock again
// 		In go Locks are not re-entering
// Solution: Centralizing locking
