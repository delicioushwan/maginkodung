package routes

import (
	"net/http"

	"github.com/delicioushwan/magickodung/configs"
	"github.com/delicioushwan/magickodung/delivery/controllers/category"
	"github.com/delicioushwan/magickodung/delivery/controllers/question"
	"github.com/delicioushwan/magickodung/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

var config = configs.GetConfig()


func Session (c echo.Context) error {
		return c.JSON(http.StatusOK, "nice")
	}

func RegisterPath(
	e *echo.Echo,
	uctrl *user.UsersController,
	qctrl *question.QuestionsController,
) {

	e.GET("/session", Session)
	e.GET("/category", category.GetCategory())

	usersGroup := e.Group("/users")
	usersGroup.POST("/signup", uctrl.Signup())
	usersGroup.POST("/login", uctrl.Login())

	questionGroup := e.Group("/questions")
	questionGroup.POST("/", qctrl.CreateQuestion())
}
