package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	testEmail := "testemail@example.com"
	code := Generate(testEmail)
	assert.Len(t, code, codeLength)
}

func TestVerifyCode(t *testing.T) {
	testEmail := "testemail@example.com"

	t.Run("valid code", func(t *testing.T) {
		code := Generate(testEmail)
		err := Verify(testEmail, code)
		assert.NoError(t, err)
	})

	t.Run("invalid code", func(t *testing.T) {
		Generate(testEmail)
		err := Verify(testEmail, "123456")
		assert.ErrorIs(t, err, ErrCodeInvalid)
	})

	t.Run("expired code", func(t *testing.T) {
		Generate(testEmail)
		codeMap.Delete(testEmail)
		err := Verify(testEmail, "123456")
		assert.ErrorIs(t, err, ErrCodeInvalid)
	})
}
