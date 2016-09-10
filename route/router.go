package route

import (
	"github.com/ken-aio/pictwitter-go/api"
	"github.com/ken-aio/pictwitter-go/db"
	"github.com/ken-aio/pictwitter-go/handler"
	myMw "github.com/ken-aio/pictwitter-go/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"os"
)

func Init() *echo.Echo {

	e := echo.New()

	if os.Getenv("ECHO_ENV") != "production" {
		e.Debug()
	}

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/fonts", api.GetFonts())
		v1.GET("/image", api.GetImage())
	}
	return e
}
