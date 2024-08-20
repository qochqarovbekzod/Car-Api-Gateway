package handler

import (
	"api_gateway/generated/auth"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// @Summary Get User Profile
// @Security BearerAuth
// @Description User
// @Tags Users
// @Param user_id path string true "ID of the user"
// @Success 200 {object} auth.UserProfileResponse
// @Failure 400 {object} string "Invalid token"
// @Failure 500 {object} string "error while ready from server"
// @Router /api/v1/users/profile/{user_id} [get]
func (h *Handler) GetUserProfile(ctx *gin.Context) {
	h.Logger.Info("GetUserProfile called with user profile information from database")
	id:=ctx.Param("user_id")
	
	// ctx.JSON(200, auth.UserProfileResponse{
	// 	Email: "string",
	// 	Password: "string",
	// 	Id: cast.ToString(id),
    //     Role: "string",
	// 	LastName: "string",
	// 	FirstName: "string",
    //     CreatedAt: "string",
    //     UpdatedAt: "string",
	// })
	// if true {
	// 	return
	// }

	resp,err:=h.Auth.GetUserProfile(ctx, &auth.Id{Id: id})

	fmt.Println(err)
	if err!=nil{
        h.Logger.Error("GetUserProfile failed", "error", err)
        ctx.JSON(500, gin.H{"error": "internal server error"})
        return
    }

	h.Logger.Info("GetUserProfile success")
	ctx.JSON(200, resp)
}

// @Summary Update user profile
// @Security BearerAuth
// Description you can update profile
// @Tags Users
// @Param user_id path string true "user_id"
// @Param userinfo body auth.UpdateUserProfileRequest true "info"
// @Success 200 {object} auth.UserProfileResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router  /api/v1/users/profile/{user_id} [put]
func (h *Handler) UpdateUserProfile(ctx *gin.Context) {
	h.Logger.Info("UpdateUserProfile called with updated user profile information")
    id:=ctx.Param("user_id")
	// ctx.JSON(200, auth.UserProfileResponse{
	// 	Email: "string",
	// 	Password: "string",
	// 	Id: cast.ToString(id),
    //     Role: "string",
	// 	LastName: "string",
	// 	FirstName: "string",
    //     CreatedAt: "string",
    //     UpdatedAt: "string",
	// })
	// if true {
	// 	return
	// }
   fmt.Println(id)
    var req auth.UpdateUserProfileRequest
    if err:=ctx.ShouldBindJSON(&req);err!=nil{
        h.Logger.Error("UpdateUserProfile failed to bind JSON", "error", err)
        ctx.JSON(400, gin.H{"error": "invalid request body"})
        return
    }

    req.Id = cast.ToString(id)
    resp,err:=h.Auth.UpdateUserProfile(ctx, &req)

    if err!=nil{
        h.Logger.Error(" UpdateUserProfile failed to update user profile", "error", err)
		ctx.JSON(500, gin.H{"error": "internal server error"})
	}
	h.Logger.Info("UpdateUserProfile success")
	ctx.JSON(200, resp)
}