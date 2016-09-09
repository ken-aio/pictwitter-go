package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/ken-aio/pictwitter-go/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetFonts() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		fonts := new(model.Fonts)
		if err = fonts.Load(tx); err != nil {
			logrus.Debug(err)
			return c.JSON(fasthttp.StatusOK, []model.Font{})
		}

		return c.JSON(fasthttp.StatusOK, fonts)
	}
}
