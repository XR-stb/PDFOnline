package static

import (
	"testing"

	"backend/pkg/config"

	"github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	config.Cfg = &config.Config{}
	config.Cfg.StaticDir = t.TempDir()

	err := Init()
	assert.NoError(t, err)

	assert.DirExists(t, PdfDir)
	assert.DirExists(t, CoverDir)
}
