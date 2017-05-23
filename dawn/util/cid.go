package util

import (
	"fmt"
	"sync/atomic"
)

// AtomicInt64 provides atomic int64 type.
type AtomicInt64 int64

// NewAtomicInt64 returns an atomic int64 type.
func NewAtomicInt64(initialValue int64) *AtomicInt64 {
	a := AtomicInt64(initialValue)
	return &a
}

// Get returns the value of int64 atomically.
func (a *AtomicInt64) Get() int64 {
	return int64(*a)
}

// CompareAndSet compares int64 with expected value, if equals as expected
// then sets the updated value, this operation performs atomically.
func (a *AtomicInt64) CompareAndSet(expect, update int64) bool {
	return atomic.CompareAndSwapInt64((*int64)(a), expect, update)
}

// GetAndIncrement gets the old value and then increment by 1, this operation
// performs atomically.
func (a *AtomicInt64) Increment() int64 {
	for {
		current := a.Get()
		next := current + 1
		if a.CompareAndSet(current, next) {
			return current
		}
	}

}

func (a *AtomicInt64) String() string {
	return fmt.Sprintf("%d", a.Get())
}
