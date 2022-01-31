package controllercategory

import (
	"net/http"
	"strconv"

	"github.com/arfan21/golang-tokobelanja/helper"
	"github.com/arfan21/golang-tokobelanja/model/modelcategory"
	"github.com/arfan21/golang-tokobelanja/service/servicecategory"
	"github.com/gin-gonic/gin"
)

type ControllerCategory interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv servicecategory.ServiceCategory
}

func New(srv servicecategory.ServiceCategory) ControllerCategory {
	return &controller{srv: srv}
}

func (c *controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	param, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	err = c.srv.Delete(param)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, gin.H{"message": "Category has been successfully deleted"}, nil))
}

func (c *controller) Update(ctx *gin.Context) {
	request := new(modelcategory.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	id := ctx.Param("id")
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	request.ID = uint(idUint64)
	update, err := c.srv.Update(*request)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, update, nil))
}

func (c *controller) GetAll(ctx *gin.Context) {
	resp, err := c.srv.GetAll()

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, resp, nil))
}

func (c *controller) Create(ctx *gin.Context) {
	request := new(modelcategory.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	create, err := c.srv.Create(*request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, create, nil))
}
