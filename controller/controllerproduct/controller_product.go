package controllerproduct

import (
	"net/http"
	"strconv"

	"github.com/arfan21/golang-tokobelanja/helper"
	"github.com/arfan21/golang-tokobelanja/model/modelproduct"
	"github.com/arfan21/golang-tokobelanja/service/serviceproduct"
	"github.com/gin-gonic/gin"
)

type ControllerProduct interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv serviceproduct.ServiceProduct
}

func New(srv serviceproduct.ServiceProduct) ControllerProduct {
	return &controller{srv: srv}
}

func (c *controller) Create(ctx *gin.Context) {
	req := modelproduct.Request{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	res, err := c.srv.Create(req)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, res, nil))
}

func (c *controller) GetAll(ctx *gin.Context) {
	res, err := c.srv.GetAll()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, res, nil))
}

func (c *controller) Update(ctx *gin.Context) {
	req := modelproduct.Request{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	id := ctx.Param("productID")
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	req.ID = uint(idUint64)

	res, err := c.srv.Update(req)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, res, nil))
}

func (c *controller) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("productID")
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	err = c.srv.DeleteByID(uint(idUint64))
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, gin.H{"message": "Product has been successfully deleted"}, nil))
}
