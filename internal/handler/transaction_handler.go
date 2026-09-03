package handler

import (
	"online-ticketing/internal/response"
	"online-ticketing/internal/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	trs *service.TransactionService
}

func NewTransactionHandler(trs *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		trs: trs,
	}
}

func (t *TransactionHandler) GetTransaction(ctx *gin.Context) {
	val, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}
	userId, ok := val.(string)
	if !ok {
		ctx.JSON(500, gin.H{
			"status":  "error",
			"message": "Internal server error",
		})
		return
	}

	result, err := t.trs.FetchTransactionByUserId(userId)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	response.ResponseSuccess(ctx, 200, "Data retrieved successfully", result)
}

func (t *TransactionHandler) CancelTransaction(ctx *gin.Context) {
	bookingCode := ctx.Param("bookingCode")
	val, exist := ctx.Get("userId")
	if !exist {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}
	userId, ok := val.(string)
	if !ok {
		ctx.JSON(500, gin.H{
			"status":  "error",
			"message": "Internal server error",
		})
		return
	}

	err := t.trs.CancelTransaction(bookingCode, userId)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	response.ResponseSuccess(ctx, 200, "Transaction cancelled successfully", nil)
}
