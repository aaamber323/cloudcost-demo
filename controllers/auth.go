package controllers

import (
	"cloudcost-demo/controllers/base"
	"cloudcost-demo/dao"
	"cloudcost-demo/model"
	"cloudcost-demo/pkg/jwtauth"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	auth := model.Auth{}

	if err := c.ShouldBindJSON(&auth); err != nil {
		err = fmt.Errorf("invalid parames error(%v)", err)
		base.ReturnError(c, base.INVALID_PARAMS, err)
		return
	}
	if err := dao.CheckAuth(auth.AccessKey, auth.SecretKey); err != nil {
		err = fmt.Errorf("user has no perssions error(%v)", err)
		base.ReturnError(c, base.ERROR_AUTH_TOKEN_USER, err)
	}

	xToken, err := jwtauth.GenToken(auth.AccessKey, auth.SecretKey)
	if err != nil {
		err := fmt.Errorf("token auth failed error(%v)", err)
		base.ReturnError(c, base.ERROR_AUTH_CHECK_TOKEN_FAIL, err)
	}
	base.ReturnJSON(c, xToken)
}
