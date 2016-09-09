package api

import (
	"net/http"

	. "gopkg.in/check.v1"
)

func (s *ApiSuite) Test_GetFonts_レスポンスが返ってくる場合は複数レコード存在すること(c *C) {
	expected := `[{"id":1,"name":"font1","download_url":"font1_url","created_at":"2016-09-06T14:00:00Z","updated_at":"2016-09-06T14:00:00Z"},{"id":2,"name":"font2","download_url":"font2_url","created_at":"2014-07-26T12:00:00Z","updated_at":"2014-07-26T12:00:00Z"}]`
	context, rec := buildGetContext(c)
	context.SetPath("/api/v1/fonts")
	tx, code, resp := requestAPI(c, GetFonts(), context, rec)
	tx.Rollback()

	c.Assert(http.StatusOK, Equals, code)
	c.Assert(expected, Equals, resp)
}

func (s *ApiSuite) Test_GetFonts_カラレスポンスの場合はカラのレスポンスが返ってくること(c *C) {
	expected := `[]`
	context, rec := buildGetContext(c)
	context.SetPath("/api/v1/fonts")
	tx := getTx(c)
	tx.DeleteFrom("fonts").Exec()
	code, resp := requestAPIWithTx(c, GetFonts(), context, rec, tx)
	tx.Rollback()

	c.Assert(http.StatusOK, Equals, code)
	c.Assert(expected, Equals, resp)
}
