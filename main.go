package wxpayapi

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

}
