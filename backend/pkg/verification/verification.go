package verification

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"backend/pkg/memcache"
)

const (
	codeExpire = time.Second * 600
	codeLength = 6
	codeSeed   = "0123456789"
)

var (
	codeMap = new(memcache.MemCache)

	ErrCodeInvalid = errors.New("verification code invalid or expired")
)

func GenerateCode(email string) string {
	code := randomCode(codeLength)
	codeMap.SetWithExpire(email, code, codeExpire)
	return code
}

func VerifyCode(email, code string) error {
	if v, ok := codeMap.Get(email); ok && v == code {
		codeMap.Delete(email)
		return nil
	}
	return ErrCodeInvalid
}

func randomCode(n int) string {
	code := ""
	for i := 0; i < n; i++ {
		code += strconv.Itoa(rand.Intn(len(codeSeed)))
	}
	return code
}
