package handlers

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

const (
	APIEndpointName = "/api"
)

type Controller interface {
	RegisterHandlers(group *routing.RouteGroup)
}

func Register(router *routing.Router, controllers ...Controller) {
	api := router.Group(APIEndpointName)
	for _, c := range controllers {
		c.RegisterHandlers(api)
	}
}
