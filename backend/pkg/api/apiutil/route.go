package apiutil

import (
	"backend/pkg/user/role"
	"github.com/gin-gonic/gin"
)

type APIRoutes interface {
	Routes() []Route
}

type Route struct {
	Method    string
	Pattern   string
	AuthLevel role.Role
	Hooks     gin.HandlersChain
	Handler   gin.HandlerFunc
}

func AddRoutes(r *gin.Engine, api APIRoutes) {
	routes := api.Routes()
	if routes == nil {
		panic("invalid api routes")
	}

	for _, route := range routes {
		handlers := make(gin.HandlersChain, 0, len(route.Hooks)+1)
		if route.Hooks != nil {
			handlers = append(handlers, route.Hooks...)
		}
		handlers = append(handlers, route.Handler)

		if route.Method != "" {
			r.Handle(route.Method, route.Pattern, handlers...)
		} else {
			r.Any(route.Pattern, handlers...)
		}
	}
}
