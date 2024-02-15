package cache

type Cache interface {
	AcquireLock(key string) error
	SetLock(key string) error
	ReleaseLock(key string) error
}
