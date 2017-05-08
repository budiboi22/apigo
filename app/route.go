package app

import (
	"fmt"

	"github.com/budiboi22/apigo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
	conf "github.com/spf13/viper"
)

func registerRoute(e *echo.Echo) {
	e.GET("/users", handler.GetUserList)
	e.GET("users/:id", handler.GetUserByID)
}

func setupRoute(err error) error {
	if err != nil {
		return err
	}

	e := echo.New()
	e.Debug = conf.GetBool("app.debug")

	//TODO add middleware here

	registerRoute(e)

	err = fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("http server failed to start."))
	if port := conf.GetString("app.port"); port != "" {
		err = e.Start(":" + port)
	}

	return err
}
