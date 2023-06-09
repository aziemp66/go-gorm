package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AllUsers(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "All Users Endpoint Hit")
}

func NewUser(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "New User Endpoint Hit")
}

func DeleteUser(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Delete User Endpoint Hit")
}

func UpdateUser(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Update User Endpoint Hit")
}
