package handler

import "github.com/labstack/echo"

func GetAllUsers(ctx echo.Context) error {
	return ctx.JSON(200, echo.Map{
		"count": 1,
		"total": 1,
		"data":  "All users",
	})
}
