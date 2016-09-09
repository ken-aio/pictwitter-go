package api

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ken-aio/pictwitter-go/db"
	"github.com/ken-aio/pictwitter-go/middleware"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	. "gopkg.in/check.v1"
)

type ApiSuite struct{}

func (s *ApiSuite) SetUpSuite(c *C) {
	// テストスイート全体での初期化処理を記述する
}

func (s *ApiSuite) SetUpTest(c *C) {
	// 個々のテストごとに実行される初期化時の共通処理を記述する
}

func (s *ApiSuite) TearDownSuite(c *C) {
	// テストスイート全体での終了後処理を記述する
}

func (s *ApiSuite) TearDownTest(c *C) {
	// 個々のテストごとに実行される終了後の共通処理を記述する
}

// GET用
func buildGetContext(c *C) (echo.Context, *httptest.ResponseRecorder) {
	return buildContext(c, echo.GET, "dummy", "")
}

// 主にPOST用
func buildContext(c *C, method string, url string, params string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req, err := http.NewRequest(method, url, strings.NewReader(params))
	c.Assert(err, IsNil)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	return context, rec
}

// 自動トランザクション
func requestAPI(c *C, handler echo.HandlerFunc, context echo.Context, rec *httptest.ResponseRecorder) (*dbr.Tx, int, string) {
	tx, err := db.InitTest().Begin()
	c.Assert(err, IsNil)

	context.Set(middleware.TxKey, tx)
	err = handler(context)
	c.Assert(err, IsNil)

	return tx, rec.Code, rec.Body.String()
}

// 手動トランザクション
func requestAPIWithTx(c *C, handler echo.HandlerFunc, context echo.Context, rec *httptest.ResponseRecorder, tx *dbr.Tx) (int, string) {
	context.Set(middleware.TxKey, tx)
	err := handler(context)
	c.Assert(err, IsNil)

	return rec.Code, rec.Body.String()
}

// Tx取得
func getTx(c *C) (tx *dbr.Tx) {
	tx, err := db.InitTest().Begin()
	c.Assert(err, IsNil)
	return
}
