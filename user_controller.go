package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func userPost(c *gin.Context) {

	user := &user{
		username:    c.PostForm("username"),
		email:       c.PostForm("email"),
		password:    myHelper.HashPassword(c.PostForm("password"), c.PostForm("username")),
		dateCreated: time.Now(),
	}

	user, err := user.Create()
	if err != nil {
		c.JSON(500, err.Error())
		c.Abort()
	}

	c.JSON(200, "{user: franco}")
}
