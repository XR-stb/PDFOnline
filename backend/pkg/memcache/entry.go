package memcache

import (
	"time"
)

type entry struct {
	k      string
	v      any
	expire time.Time
}

func (e *entry) expired() bool {
	if e.expire.IsZero() {
		return false
	}
	return e.expire.Before(time.Now())
}
