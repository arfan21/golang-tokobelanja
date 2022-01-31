package controllertransactionhistory

import (
	"github.com/arfan21/golang-tokobelanja/helper"
	"github.com/arfan21/golang-tokobelanja/model/modeltransactionhistory"
	"github.com/arfan21/golang-tokobelanja/service/servicetransactionhistory"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerTransactionHistory interface {
	CreateTransaction(ctx *gin.Context)
	GetTransactions(ctx *gin.Context)
	GetTransactionByUserHistories(ctx *gin.Context)
}

type Controller struct {
	srv servicetransactionhistory.ServiceTransactionHistory
}

func (c *Controller) GetTransactions(ctx *gin.Context) {
	transactions, err := c.srv.GetTransactionHistories()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, transactions, nil))
}

func (c *Controller) GetTransactionByUserHistories(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	userHistories, err := c.srv.GetTransactionHistory(userID)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, userHistories, nil))
}

func (c *Controller) CreateTransaction(ctx *gin.Context) {
	request := modeltransactionhistory.RequestTransaction{}
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, "Invalid request body", err))
		return
	}
	uID := ctx.MustGet("user_id").(uint)
	request.UserID = uID
	response, err := c.srv.CreateTransactionHistory(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), "Failed to create transaction history", err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

func New(srv servicetransactionhistory.ServiceTransactionHistory) ControllerTransactionHistory {
	return &Controller{srv: srv}
}
