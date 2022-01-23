package server

import (
	"github.com/arfan21/golang-tokobelanja/controller/controlleruser"
	"github.com/arfan21/golang-tokobelanja/middleware"
	"github.com/arfan21/golang-tokobelanja/repository/repositoryuser"
	"github.com/arfan21/golang-tokobelanja/service/serviceuser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/topup", middleware.Authorization, ctrlUser.Update)
}
