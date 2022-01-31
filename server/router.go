package server

import (
	"github.com/arfan21/golang-tokobelanja/controller/controllercategory"
	"github.com/arfan21/golang-tokobelanja/controller/controllerproduct"
	"github.com/arfan21/golang-tokobelanja/controller/controllertransactionhistory"
	"github.com/arfan21/golang-tokobelanja/controller/controlleruser"
	"github.com/arfan21/golang-tokobelanja/middleware"
	"github.com/arfan21/golang-tokobelanja/repository/repositorycategory"
	"github.com/arfan21/golang-tokobelanja/repository/repositoryproduct"
	"github.com/arfan21/golang-tokobelanja/repository/repositorytransactionhistory"
	"github.com/arfan21/golang-tokobelanja/repository/repositoryuser"
	"github.com/arfan21/golang-tokobelanja/service/servicecategory"
	"github.com/arfan21/golang-tokobelanja/service/serviceproduct"
	"github.com/arfan21/golang-tokobelanja/service/servicetransactionhistory"
	"github.com/arfan21/golang-tokobelanja/service/serviceuser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	// route user
	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PATCH("/topup", middleware.Authorization, ctrlUser.Update)

	// route category
	repoCategory := repositorycategory.New(db)
	srvCategory := servicecategory.New(repoCategory)
	ctrlCategory := controllercategory.New(srvCategory)

	routeCategory := r.Group("/categories")

	routeCategory.POST("", middleware.Authorization, middleware.AuthorizationAdmin, ctrlCategory.Create)
	routeCategory.GET("", middleware.Authorization, middleware.AuthorizationAdmin, ctrlCategory.GetAll)
	routeCategory.PATCH("/:id", middleware.Authorization, middleware.AuthorizationAdmin, ctrlCategory.Update)
	routeCategory.DELETE("/:id", middleware.Authorization, middleware.AuthorizationAdmin, ctrlCategory.Delete)

	// route product
	repoProduct := repositoryproduct.New(db)
	srvProduct := serviceproduct.New(repoProduct, repoCategory)
	ctrlProduct := controllerproduct.New(srvProduct)

	routeProduct := r.Group("/products")
	routeProduct.POST("", middleware.Authorization, middleware.AuthorizationAdmin, ctrlProduct.Create)
	routeProduct.GET("", middleware.Authorization, ctrlProduct.GetAll)
	routeProduct.PUT("/:productID", middleware.Authorization, middleware.AuthorizationAdmin, ctrlProduct.Update)
	routeProduct.DELETE("/:productID", middleware.Authorization, middleware.AuthorizationAdmin, ctrlProduct.DeleteByID)

	// route Transaction History
	repoTransaction := repositorytransactionhistory.New(db)
	srvTransaction := servicetransactionhistory.New(repoTransaction)
	ctrlTransaction := controllertransactionhistory.New(srvTransaction)

	routeTransaction := r.Group("/transactions")
	routeTransaction.POST("", middleware.Authorization, ctrlTransaction.CreateTransaction)
	routeTransaction.GET("my-transactions", middleware.Authorization, ctrlTransaction.GetTransactionByUserHistories)
	routeTransaction.GET("user-transaction", middleware.Authorization, ctrlTransaction.GetTransactions)
}
