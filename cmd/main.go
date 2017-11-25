package main

import (
	"flag"
	wxpayapi "lemon-wxpay-api"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	appEnv   = flag.String("APP_ENV", os.Getenv("APP_ENV"), "APP_ENV")
	appId    = flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	key      = flag.String("WXPAY_KEY", os.Getenv("WXPAY_KEY"), "WXPAY_KEY")
	mchId    = flag.String("WXPAY_MCHID", os.Getenv("WXPAY_MCHID"), "WXPAY_MCHID")
	certName = flag.String("CERT_NAME", os.Getenv("CERT_NAME"), "CERT_NAME")
	certKey  = flag.String("CERT_KEY", os.Getenv("CERT_KEY"), "CERT_KEY")
	rootCa   = flag.String("ROOT_CA", os.Getenv("ROOT_CA"), "ROOT_CA")
)

func main() {
	flag.Parse()
	wxpayapi.EnvParam = &wxpayapi.EnvParamDto{
		AppEnv:   *appEnv,
		AppId:    *appId,
		Key:      *key,
		MchId:    *mchId,
		CertName: *certName,
		CertKey:  *certKey,
		RootCa:   *rootCa,
	}
	//initTest()
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
	green.POST("/pay", wxpayapi.PayGreen)
	green.POST("/query", wxpayapi.QueryGreen)
	green.POST("/reverse", wxpayapi.ReverseGreen)
	green.POST("/refund", wxpayapi.RefundGreen)
	green.POST("/refundquery", wxpayapi.RefundQueryGreen)
	green.POST("/prepay", wxpayapi.PrePayGreen)
	green.POST("/notify", wxpayapi.NotifyGreen)

}
