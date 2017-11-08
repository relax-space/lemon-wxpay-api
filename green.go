package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/relax-space/go-kit/model"

	"github.com/relax-space/lemon-wxpay"

	"github.com/labstack/echo"
)

func PayGreen(c echo.Context) error {
	reqDto := wxpay.ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account := Account()
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Pay(reqDto, customDto)
	if err != nil {
		if err.Error() == "MESSAGE_PAYING" {
			queryDto := wxpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			result, err = wxpay.LoopQuery(queryDto, customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result["out_trade_no"].(string),
				}
				_, err = wxpay.Reverse(reverseDto, customDto, 10, 10)
				return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
		}
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func QueryGreen(c echo.Context) error {
	reqDto := wxpay.ReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account := Account()
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Query(reqDto, customDto)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}
func RefundGreen(c echo.Context) error {
	return nil
}
func ReverseGreen(c echo.Context) error {
	return nil
}

func Account() greenAccount {

	appId := flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	key := flag.String("WXPAY_KEY", os.Getenv("WXPAY_KEY"), "WXPAY_KEY")
	mchId := flag.String("WXPAY_MCHID", os.Getenv("WXPAY_MCHID"), "WXPAY_MCHID")
	account := greenAccount{AppId: *appId, Key: *key, MchId: *mchId}
	return account
}

type greenAccount struct {
	AppId string
	Key   string
	MchId string
}
