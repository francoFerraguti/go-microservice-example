package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var myRouter *router

type Router interface {
	Get() *gin.Engine
	Set(router *gin.Engine)
	Init()
	Run()
	Pong()
}

type router struct {
	router *gin.Engine
}

func (r *router) Get() *gin.Engine {
	return r.router
}

func (r *router) Set(router *gin.Engine) {
	log.Println("", r)
	r.router = router
}

func (r *router) Init() {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	}))

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", r.Pong)

		v1.POST("/user", userPost)
	}

	r.Set(router)
}

func (r *router) Run() {
	log.Fatal(r.router.Run(":" + myConfiguration.Port()))
}

func (r *router) Pong(c *gin.Context) {
	c.String(200, "pong")
}
