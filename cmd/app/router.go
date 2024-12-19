package app

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-mysql/internal/app"
)

type Router struct {
	addr   string
	router *gin.Engine
}

func NewRouter(addr string) *Router {
	return &Router{
		addr:   addr,
		router: gin.Default(),
	}
}

func (r *Router) Config() {
	r.router.MaxMultipartMemory = 8 << 20 // 8 MiB

	v := r.router.Group("/v1/user")
	v.GET("", app.GetUserList)
}

func (r *Router) Run() {
	go r.router.Run(r.addr)
}
