package handler

import (
	"learning-gomod/domain/dao"
	"learning-gomod/domain/service"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	findAllUserService, errorHandlerService := service.FetchAllUsers()

	if errorHandlerService != nil {
		log.Printf("Error when get service %s", errorHandlerService)
	}

	if findAllUserService == nil {
		ctx.JSON(200, gin.H{
			"total": len(findAllUserService),
			"count": len(findAllUserService),
			"data":  []int{},
		})
	}

	ctx.JSON(200, gin.H{
		"total": len(findAllUserService),
		"count": len(findAllUserService),
		"data":  findAllUserService,
	})

}

func GetDetailUsers(ctx *gin.Context) {
	UsersParam := ctx.Param("UsersParam")
	convertUsersID, _ := strconv.Atoi(UsersParam)

	findDetailUserService, errorHandlerService := service.FetchUserByID(convertUsersID)

	if errorHandlerService != nil {
		log.Printf("Error when get service %s", errorHandlerService)
	}

	if findDetailUserService == nil {
		ctx.JSON(200, gin.H{
			"data": []int{},
		})
	}

	ctx.JSON(200, gin.H{
		"data": findDetailUserService,
	})
}

func CreateNewUsers(ctx *gin.Context) {
	var userPayload dao.CreateNewUser

	ctx.BindJSON(&userPayload)

	userPayload.CreatedAt = time.Now()
	userPayload.UpdatedAt = time.Now()

	errorHandlerService := service.SaveNewUser(userPayload)

	if errorHandlerService != nil {
		ctx.JSON(500, errorHandlerService.Error())
	}

	ctx.JSON(200, gin.H{
		"message": "User created",
		"created": userPayload,
	})
}
