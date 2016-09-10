package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/ken-aio/pictwitter-go/middleware"
	"github.com/ken-aio/pictwitter-go/model"
	"github.com/gocraft/dbr"
	"os/exec"
	"strconv"
	"strings"
)

func GetImage() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		word := c.QueryParam("word")
		if word == "" {
			logrus.Debug("word is " + word)
			return c.JSON(fasthttp.StatusOK, map[string]string{"image": ""})
		}
		fontId, err := strconv.Atoi(c.QueryParam("font_id"))
		if err != nil {
			fontId = 1
		}
		logrus.Info(fontId)
		font := model.NewFont()
		tx := c.Get(middleware.TxKey).(*dbr.Tx)
		filename := "GenEiAntiqueP-Medium.otf"
		font.LoadColumnById(tx, fontId, "file_name").Load(&filename)
		logrus.Info(filename)
		cmdstr := "convert -font assets/fonts/" + filename + " -pointsize 16 label:'" + word + "' png:- | base64"
		out, err := exec.Command("sh", "-c", cmdstr).Output()

		return c.JSON(fasthttp.StatusOK, map[string]string{"image": "data:image/jpeg;base64," + strings.TrimRight(string(out), "\n")})
	}
}
