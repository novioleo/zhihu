package controllers

import (
	"../config"
	"../models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Visitor(c *gin.Context) (*models.User, uint) {
	sess := sessions.Default(c)
	uid := sess.Get(config.Server.SessionKey)
	if uid == nil {
		return nil, 0
	}
	user := models.GetUserByID(uid.(uint))
	if user == nil {
		return nil, 0
	}
	return user, uid.(uint)
}

func VisitorID(c *gin.Context) uint {
	sess := sessions.Default(c)
	uid := sess.Get(config.Server.SessionKey)
	if uid == nil {
		return 0
	}
	return uid.(uint)
}
