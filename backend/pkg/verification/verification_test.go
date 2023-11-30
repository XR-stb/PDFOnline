package verification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	testEmail := "testemail@example.com"
	code := GenerateCode(testEmail)
	assert.Len(t, code, codeLength)
}

func TestVerifyCode(t *testing.T) {
	testEmail := "testemail@example.com"

	t.Run("valid code", func(t *testing.T) {
		code := GenerateCode(testEmail)
		err := VerifyCode(testEmail, code)
		assert.NoError(t, err)
	})

	t.Run("invalid code", func(t *testing.T) {
		GenerateCode(testEmail)
		err := VerifyCode(testEmail, "123456")
		assert.ErrorIs(t, err, ErrCodeInvalid)
	})

	t.Run("expired code", func(t *testing.T) {
		GenerateCode(testEmail)
		codeMap.Delete(testEmail)
		err := VerifyCode(testEmail, "123456")
		assert.ErrorIs(t, err, ErrCodeInvalid)
	})
}
