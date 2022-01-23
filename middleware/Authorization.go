package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/arfan21/golang-tokobelanja/constant"
	"github.com/arfan21/golang-tokobelanja/helper"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, errors.New("the request is allowed for logged in")))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	bearerToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	id, role, err := helper.ParseJwt(bearerToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, err))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_id", id)
	c.Set("role", role)
	c.Next()
}

func AuthorizationAdmin(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != constant.AdminRole {
		c.JSON(http.StatusForbidden, helper.NewResponse(http.StatusForbidden, nil, errors.New("your role is not allowed")))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	log.Println(role)
}
