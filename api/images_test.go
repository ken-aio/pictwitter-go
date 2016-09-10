package api

import (
	"net/http"
	"strings"

	. "gopkg.in/check.v1"
	"github.com/labstack/echo"
)

func (s *ApiSuite) Test_GetImage_単語を指定した場合にBase64の文字列が返ってくること(c *C) {
	expected := "base64"
	context, rec := buildContext(c, echo.GET, "/?word=test", "")
	context.SetPath("/api/v1/image")
	tx, code, resp := requestAPI(c, GetImage(), context, rec)
	tx.Rollback()

	c.Assert(http.StatusOK, Equals, code)
	c.Assert(true, Equals, strings.Contains(resp, expected))
}

func (s *ApiSuite) Test_GetImage_単語を指定しなかった場合にカラの文字列が返ってくること(c *C) {
	expected := `{"image":""}`
	context, rec := buildGetContext(c)
	context.SetPath("/api/v1/image")
	tx, code, resp := requestAPI(c, GetImage(), context, rec)
	tx.Rollback()

	c.Assert(http.StatusOK, Equals, code)
	c.Assert(expected, Equals, resp)
}
