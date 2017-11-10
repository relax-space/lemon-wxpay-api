package main

import (
	"fmt"
	"net/http"

	"github.com/relax-space/go-kit/model"

	"github.com/relax-space/lemon-wxpay"

	"github.com/labstack/echo"
)

func PayGreen(c echo.Context) error {
	reqDto := wxpay.ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
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
	reqDto := wxpay.ReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	account := Account()
	fmt.Println(account)
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	custDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertPathName,
		CertPathKey:  account.CertPathKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Refund(reqDto, custDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})

}
func ReverseGreen(c echo.Context) error {
	reqDto := wxpay.ReqReverseDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	account := Account()
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	custDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertPathName,
		CertPathKey:  account.CertPathKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Reverse(reqDto, custDto, 10, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func RefundQueryGreen(c echo.Context) error {
	reqDto := wxpay.ReqRefundQueryDto{}
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
	result, err := wxpay.RefundQuery(reqDto, customDto)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func Account() greenAccount {

	account := greenAccount{AppId: *appId, Key: *key, MchId: *mchId,
		CertPathName: *certName, CertPathKey: *certKey, RootCa: *rootCa,
	}
	return account
}

type greenAccount struct {
	AppId        string
	Key          string
	MchId        string
	CertPathName string
	CertPathKey  string
	RootCa       string
}
