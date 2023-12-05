package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"backend/pkg/api/apiutil"
	v1 "backend/pkg/api/v1"
	"backend/pkg/config"
	"backend/pkg/static"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	setMode()
	setOutput()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
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
