package handler

import (
	"learning-gomod/domain/dao"
	"learning-gomod/domain/service"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllUserHandler(ctx *gin.Context) {
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

func GetDetailUserHandler(ctx *gin.Context) {
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

func CreateNewUserHandler(ctx *gin.Context) {
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

func UpdateUserHandler(ctx *gin.Context) {
	var (
		userPayload dao.UpdateUser
		UsersParam  = ctx.Param("UsersParam")
	)

	convertUsersID, _ := strconv.Atoi(UsersParam)

	ctx.BindJSON(&userPayload)

	userPayload.CreatedAt = time.Now()
	userPayload.UpdatedAt = time.Now()

	findDetailUserService, errorHandlerService := service.FetchUserByID(convertUsersID)

	if errorHandlerService != nil {
		log.Printf("Error when get service %s", errorHandlerService)
	}

	if findDetailUserService == nil {
		ctx.JSON(200, gin.H{
			"message": "Data not found",
		})
	} else {
		updateUserProcess(userPayload, convertUsersID, ctx)
	}

}

func updateUserProcess(updateUserPayload dao.UpdateUser, convertedUserID int, ctx *gin.Context) {
	errorHandlerUpdateUserService := service.UpdateUser(convertedUserID, updateUserPayload)

	if errorHandlerUpdateUserService != nil {
		ctx.JSON(500, gin.H{
			"error": errorHandlerUpdateUserService,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message":        "User updated successfully",
			"updated_record": updateUserPayload,
		})
	}
}
