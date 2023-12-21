package data

type Lock struct {
	Key int
}

type LockerInterface interface {
	StartLock() *Lock
	Abort(lock *Lock)
	EndLock(lock *Lock)
}
