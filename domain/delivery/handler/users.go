package handler

import (
	"learning-gomod/domain/service"
	"log"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllUsers(ctx echo.Context) error {
	findAllUserService, errorHandlerService := service.FetchAllUsers()

	if errorHandlerService != nil {
		log.Printf("Error when get service %s", errorHandlerService)
	}

	if findAllUserService == nil {
		return ctx.JSON(200, echo.Map{
			"total": len(findAllUserService),
			"count": len(findAllUserService),
			"data":  []int{},
		})
	}

	return ctx.JSON(200, echo.Map{
		"total": len(findAllUserService),
		"count": len(findAllUserService),
		"data":  findAllUserService,
	})

}

func GetDetailUsers(ctx echo.Context) error {
	UsersParam := ctx.Param("UsersParam")
	convertUsersID, _ := strconv.Atoi(UsersParam)

	findDetailUserService, errorHandlerService := service.FetchUserByID(convertUsersID)

	if errorHandlerService != nil {
		log.Printf("Error when get service %s", errorHandlerService)
	}

	if findDetailUserService == nil {
		return ctx.JSON(200, echo.Map{
			"data": []int{},
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": findDetailUserService,
	})
}
