package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/relax-space/go-kit/model"

	"github.com/relax-space/lemon-wxpay-sdk"

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
	result, err := wxpay.Pay(&reqDto, &customDto)
	if err != nil {
		if err.Error() == "MESSAGE_PAYING" {
			queryDto := wxpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			result, err = wxpay.LoopQuery(&queryDto, &customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result["out_trade_no"].(string),
				}
				_, err = wxpay.Reverse(&reverseDto, &customDto, 10, 10)
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
	result, err := wxpay.Query(&reqDto, &customDto)
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
	result, err := wxpay.Refund(&reqDto, &custDto)
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
	result, err := wxpay.Reverse(&reqDto, &custDto, 10, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func RefundQueryGreen(c echo.Context) error {
	reqDto := wxpay.ReqRefundQueryDto{}
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
	result, err := wxpay.RefundQuery(&reqDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func PrePayGreen(c echo.Context) error {
	reqDto := wxpay.ReqPrePayDto{}
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
	result, err := wxpay.PrePay(&reqDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func NotifyGreen(c echo.Context) error {

	errResult := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{xml.Name{}, "FAIL", ""}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}
	xmlBody := string(body)
	if len(xmlBody) == 0 {
		return c.XML(http.StatusBadRequest, errResult)
	}
	result, err := wxpay.Notify(xmlBody)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}
	return c.XML(http.StatusOK, result)
}

func Account() greenAccount {

	account := greenAccount{AppId: envParam.AppId, Key: envParam.Key, MchId: envParam.MchId,
		CertPathName: envParam.CertName, CertPathKey: envParam.CertKey, RootCa: envParam.RootCa,
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
