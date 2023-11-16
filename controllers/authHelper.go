package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func checkUsertype(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized To Access This Resource")
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized To Access This Resource")
		return err
	}
	err = checkUsertype(c, userType)
	return err
}
