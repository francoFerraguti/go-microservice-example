package main

import (
	"github.com/francoFerraguti/go-microservice-example/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Get() *gin.Engine
	Set(router *gin.Engine)
	Init()
	Run()
}

type router struct {
	router *gin.Engine
}

func (r *router) Get() *gin.Engine {
	return r.router
}

func (r *router) Set(router *gin.Engine) {
	r.router = router
}

func (r *router) Init() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET,PUT,POST,DELETE"},
		AllowHeaders:    []string{"accept,x-access-token,content-type,authorization"},
	}))

	v1 := router.Group("/v1")
	{
		v1.POST("/user", user.Post)
	}

	r.Set(router)
}

func (r *router) Run() {
	r.router.Run(":80")
}
