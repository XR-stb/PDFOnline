package testutil

import (
	"backend/pkg/user/role"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/pkg/database"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestDB(t *testing.T) *gorm.DB {
	dsn := "file:pdfserver.db?mode=memory"
	mockDB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)
	err = database.AutoMigrate(mockDB)
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

func StringPtr(s string) *string {
	return &s
}

func RolePtr(r role.Role) *role.Role {
	return &r
}
