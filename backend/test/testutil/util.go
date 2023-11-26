package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/pkg/database"
)

func TestDB(t *testing.T) *gorm.DB {
	dsn := "file:furion.db?mode=memory"
	mockDB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)
	err = database.AutoMigrate(mockDB)
	require.NoError(t, err)
	// SQLite: enable foreign key check
	err = mockDB.Exec("PRAGMA foreign_keys = ON").Error
	require.NoError(t, err)
	return mockDB
}

func NewRequest(t *testing.T, method, url string, obj ...any) *http.Request {
	var req *http.Request
	if len(obj) == 0 {
		req = httptest.NewRequest(method, url, nil)
	} else {
		b, err := json.Marshal(obj[0])
		assert.NoError(t, err)
		req = httptest.NewRequest(method, url, bytes.NewReader(b))
	}

	req.Header.Set("Content-Type", "application/json")

	return req
}
