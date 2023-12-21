package utils

import (
	"sync"
	"tokbel/services/data"

	"gorm.io/gorm"
)

type Locker struct {
	db      *gorm.DB
	mu      sync.Mutex
	counter int
	locks   map[int]*gorm.DB
}

func NewLocker(db *gorm.DB) *Locker {
	return &Locker{
		db:      db,
		mu:      sync.Mutex{},
		counter: 0,
		locks:   make(map[int]*gorm.DB),
	}
}

func (locker *Locker) StartLock() *data.Lock {
	locker.mu.Lock()
	defer locker.mu.Unlock()
	locker.counter++
	locker.locks[locker.counter] = locker.db.Begin()
	return &data.Lock{Key: locker.counter}
}

func (locker *Locker) GetLock(lock *data.Lock) *gorm.DB {
	return locker.locks[lock.Key]
}

func (locker *Locker) Abort(lock *data.Lock) {
	locker.mu.Lock()
	defer locker.mu.Unlock()
	key := lock.Key
	tx := locker.locks[key]
	tx.Rollback()
	delete(locker.locks, key)
}

func (locker *Locker) EndLock(lock *data.Lock) {
	locker.mu.Lock()
	defer locker.mu.Unlock()
	key := lock.Key
	tx := locker.locks[key]
	tx.Commit()
	delete(locker.locks, key)
}
