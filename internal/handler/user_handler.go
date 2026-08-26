package handler

import (
	"online-ticketing/internal/model"
	"online-ticketing/internal/response"
	"online-ticketing/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	user := model.RegisterRequest{}

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	result, err := u.us.Register(user)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	response.ResponseSuccess(ctx, 201, "User create successfully", result)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	req := model.LoginRequest{}

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	result, err := u.us.Login(req)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	response.ResponseSuccess(ctx, 200, "Successfully login", result)
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	result, err := u.us.FetchAllUsers()
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	response.ResponseSuccess(ctx, 200, "Data retrieved successfully", result)
}
