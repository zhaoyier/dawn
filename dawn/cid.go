package dawn

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

// Set sets the value of int64 atomically.
func (a *AtomicInt64) Set(newValue int64) {
	atomic.StoreInt64((*int64)(a), newValue)
}

// GetAndSet sets new value and returns the old atomically.
func (a *AtomicInt64) GetAndSet(newValue int64) int64 {
	for {
		current := a.Get()
		if a.CompareAndSet(current, newValue) {
			return current
		}
	}
}

// CompareAndSet compares int64 with expected value, if equals as expected
// then sets the updated value, this operation performs atomically.
func (a *AtomicInt64) CompareAndSet(expect, update int64) bool {
	return atomic.CompareAndSwapInt64((*int64)(a), expect, update)
}

// GetAndIncrement gets the old value and then increment by 1, this operation
// performs atomically.
func (a *AtomicInt64) GetAndIncrement() int64 {
	for {
		current := a.Get()
		next := current + 1
		if a.CompareAndSet(current, next) {
			return current
		}
	}

}

// GetAndDecrement gets the old value and then decrement by 1, this operation
// performs atomically.
func (a *AtomicInt64) GetAndDecrement() int64 {
	for {
		current := a.Get()
		next := current - 1
		if a.CompareAndSet(current, next) {
			return current
		}
	}
}

// GetAndAdd gets the old value and then add by delta, this operation
// performs atomically.
func (a *AtomicInt64) GetAndAdd(delta int64) int64 {
	for {
		current := a.Get()
		next := current + delta
		if a.CompareAndSet(current, next) {
			return current
		}
	}
}

// IncrementAndGet increments the value by 1 and then gets the value, this
// operation performs atomically.
func (a *AtomicInt64) IncrementAndGet() int64 {
	for {
		current := a.Get()
		next := current + 1
		if a.CompareAndSet(current, next) {
			return next
		}
	}
}

// DecrementAndGet decrements the value by 1 and then gets the value, this
// operation performs atomically.
func (a *AtomicInt64) DecrementAndGet() int64 {
	for {
		current := a.Get()
		next := current - 1
		if a.CompareAndSet(current, next) {
			return next
		}
	}
}

// AddAndGet adds the value by delta and then gets the value, this operation
// performs atomically.
func (a *AtomicInt64) AddAndGet(delta int64) int64 {
	for {
		current := a.Get()
		next := current + delta
		if a.CompareAndSet(current, next) {
			return next
		}
	}
}

func (a *AtomicInt64) String() string {
	return fmt.Sprintf("%d", a.Get())
}
