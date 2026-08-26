package handler

import (
	"online-ticketing/internal/model"
	"online-ticketing/internal/response"
	"online-ticketing/internal/service"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ts *service.TicketService
}

func NewTicketHandler(ts *service.TicketService) *TicketHandler {
	return &TicketHandler{
		ts: ts,
	}
}

func (t *TicketHandler) GetTickets(ctx *gin.Context) {
	result, err := t.ts.FetchAllTickets()
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	response.ResponseSuccess(ctx, 200, "Data retrieved successfully", result)
}

func (t *TicketHandler) CreateTicket(ctx *gin.Context) {
	ticket := model.Ticket{}

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	result, err := t.ts.CreateTicket(ticket)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	response.ResponseSuccess(ctx, 200, "Data retrieved successfully", result)
}

func (t *TicketHandler) BuyTicket(ctx *gin.Context) {
	ticket := model.BuyTicketRequest{}
	val, exist := ctx.Get("userId")

	if !exist {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"message": "Unauthorized",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid request body",
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
	err := t.ts.BuyTicket(ticket, userId)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	response.ResponseSuccess(ctx, 200, "Ticket bought successfully", nil)
}
