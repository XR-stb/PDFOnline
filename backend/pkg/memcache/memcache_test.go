package memcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemCache_Set(t *testing.T) {
	m := new(MemCache)
	m.Set("KEY", "VALUE")
	v, ok := m.data["KEY"]
	assert.True(t, ok)
	assert.Equal(t, "VALUE", v.v)
}

func TestMemCache_SetWithExpire(t *testing.T) {
	m := new(MemCache)
	m.SetWithExpire("KEY", "VALUE", time.Second)
	v, ok := m.data["KEY"]
	assert.True(t, ok)
	assert.Equal(t, "VALUE", v.v)

	time.Sleep(time.Second)
	assert.True(t, v.expire.Before(time.Now()))
}

func TestMemCache_Get(t *testing.T) {
	m := new(MemCache)
	m.Set("KEY", "VALUE")
	v, ok := m.Get("KEY")
	assert.True(t, ok)
	assert.Equal(t, "VALUE", v)
}

func TestMemCache_Delete(t *testing.T) {
	m := new(MemCache)
	m.Set("KEY", "VALUE")
	m.Delete("KEY")
	_, ok := m.data["KEY"]
	assert.False(t, ok)
}
