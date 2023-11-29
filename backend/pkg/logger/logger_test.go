package logger

import (
	"path"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"backend/pkg/config"
)

func TestInit(t *testing.T) {
	t.Run("set output", func(t *testing.T) {
		dir := t.TempDir()
		config.Cfg = &config.Config{Default: config.Default{Debug: true, LogDir: dir}}

		err := Init()
		assert.NoError(t, err)

		assert.EqualValues(t, logrus.DebugLevel, logrus.GetLevel())
		assert.FileExists(t, path.Join(dir, "backend.log"))
	})

	t.Run("not set output", func(t *testing.T) {
		config.Cfg = &config.Config{}

		err := Init()
		assert.NoError(t, err)

		assert.EqualValues(t, logrus.InfoLevel, logrus.GetLevel())
	})
}
