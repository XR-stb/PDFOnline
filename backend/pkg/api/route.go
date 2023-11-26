package api

import (
	"backend/pkg/api/apiutil"
	v1 "backend/pkg/api/v1"
	"backend/pkg/config"
	"backend/pkg/static"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	setMode()
	setOutput()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // 允许所有来源
		AllowMethods: []string{"*"}, // 允许所有方法
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	r.Static(apiutil.StaticRootPdf, static.PdfDir)
	r.Static(apiutil.StaticRootCover, static.CoverDir)

	apiutil.AddRoutes(r, v1.PdfAPI{})
	apiutil.AddRoutes(r, v1.UserAPI{})

	return r
}

func setMode() {
	if config.Debug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func setOutput() {
	gin.DefaultWriter = logrus.StandardLogger().Writer()
	gin.DefaultErrorWriter = logrus.StandardLogger().WriterLevel(logrus.ErrorLevel)
}
