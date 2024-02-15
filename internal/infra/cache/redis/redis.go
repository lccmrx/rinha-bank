package redis

import (
	"fmt"
	"sync"

	"github.com/lccmrx/rinha-bank/internal/infra/cache"
	"github.com/lccmrx/rinha-bank/internal/infra/config"
	"gopkg.in/redis.v5"
)

type Redis struct {
	client *redis.Client
	locker sync.Mutex
}

var _ cache.Cache = (*Redis)(nil)

func New(cfg *config.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port),
		DB:       cfg.Cache.DB,
		Password: cfg.Cache.Pass,
	})
	return &Redis{client: client}
}

func (r *Redis) AcquireLock(key string) error {
	for {
		s, _ := r.client.Get(key).Result()
		if s == "" {
			break
		}
	}

	return r.SetLock(key)
}

func (r *Redis) SetLock(key string) error {
	r.locker.Lock()
	defer r.locker.Unlock()

	ok, err := r.client.SetNX(key, key, 0).Result()
	if !ok || err != nil {
		return err
	}

	return nil
}

func (r *Redis) ReleaseLock(key string) error {
	r.locker.Lock()
	defer r.locker.Unlock()

	for {
		err := r.client.Watch(func(tx *redis.Tx) error {
			result, err := tx.Get(key).Result()
			if err != nil {
				if err == redis.Nil {
					return nil
				}
				return err
			}

			if result == key {
				err = tx.Del(key).Err()
				if err != nil {
					return err
				}
			}

			return nil
		}, key)
		if err != nil {
			continue
		}

		return nil
	}
}
