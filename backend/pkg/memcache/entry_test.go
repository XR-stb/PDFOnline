package memcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_entry_expired(t *testing.T) {
	testEntry := &entry{
		expire: time.Now().Add(-time.Second),
	}

	assert.True(t, testEntry.expired())
}
