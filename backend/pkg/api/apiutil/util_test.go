package apiutil

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Req struct {
	Username string `json:"username" binding:"required,min=6,max=36" required:"username is required" min:"username is too short, min 6 chars" max:"username is too long, max 36 chars"`
	Password string `json:"password" binding:"required,min=6,max=36" required:"password is required" min:"password is too short, min 6 chars" max:"password is too long, max 36 chars"`
}

func TestShouldBindJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	b, err := json.Marshal(Req{
		Username: "user",
	})
	assert.NoError(t, err)
	c.Request = httptest.NewRequest("POST", "/", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	err = ShouldBind(c, &Req{})
	assert.Equal(t, err.Error(), "username is too short, min 6 chars\npassword is required")
}
