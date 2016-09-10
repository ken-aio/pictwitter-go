package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"os/exec"
	"strings"
)

func GetImage() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		word := c.QueryParam("word")
		if word == "" {
			logrus.Debug("word is " + word)
			return c.JSON(fasthttp.StatusOK, map[string]string{"image": ""})
		}
		cmdstr := "convert -font assets/fonts/GenEiAntiqueP-Medium.otf -pointsize 16 label:'" + word + "' png:- | base64"
		out, err := exec.Command("sh", "-c", cmdstr).Output()

		return c.JSON(fasthttp.StatusOK, map[string]string{"image": "data:image/jpeg;base64," + strings.TrimRight(string(out), "\n")})
	}
}
