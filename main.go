package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	appId    = flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	key      = flag.String("WXPAY_KEY", os.Getenv("WXPAY_KEY"), "WXPAY_KEY")
	mchId    = flag.String("WXPAY_MCHID", os.Getenv("WXPAY_MCHID"), "WXPAY_MCHID")
	certName = flag.String("CERT_NAME", os.Getenv("CERT_NAME"), "CERT_NAME")
	certKey  = flag.String("CERT_KEY", os.Getenv("CERT_KEY"), "CERT_KEY")
	rootCa   = flag.String("ROOT_CA", os.Getenv("ROOT_CA"), "ROOT_CA")
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	RegApi(e)
	e.Start(":5000")
}
func RegApi(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello wx pay")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/v1")
	green := v1.Group("/green")
	green.POST("/pay", PayGreen)
	green.POST("/query", QueryGreen)
	green.POST("/reverse", ReverseGreen)
	green.POST("/refund", RefundGreen)
	green.POST("/refundquery", RefundQueryGreen)
	green.POST("/prepay", PrePayGreen)
	green.POST("/notify", NotifyGreen)

}
