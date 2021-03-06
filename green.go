package wxpayapi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/relax-space/go-kit/model"

	"github.com/relax-space/lemon-wxpay-sdk"

	"github.com/labstack/echo"
)

var (
	EnvParam = &EnvParamDto{}
)

type EnvParamDto struct {
	AppEnv   string
	AppId    string
	Key      string
	MchId    string
	CertName string
	CertKey  string
	RootCa   string
}

func PayGreen(c echo.Context) error {
	reqDto := wxpay.ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}

	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
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
				return c.JSON(http.StatusOK, SuccessResult(result))
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result["out_trade_no"].(string),
				}
				_, err = wxpay.Reverse(&reverseDto, &customDto, 10, 10)
				return c.JSON(http.StatusInternalServerError, ErrorResult(err.Error()))
			}
		} else {
			return c.JSON(http.StatusInternalServerError, ErrorResult(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

func QueryGreen(c echo.Context) error {
	reqDto := wxpay.ReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}

	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Query(&reqDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusOK, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}
func RefundGreen(c echo.Context) error {
	reqDto := wxpay.ReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account()
	fmt.Println(account)
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
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
		return c.JSON(http.StatusInternalServerError, ErrorResult(err.Error()))

	}
	return c.JSON(http.StatusOK, SuccessResult(result))

}
func ReverseGreen(c echo.Context) error {
	reqDto := wxpay.ReqReverseDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
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
		return c.JSON(http.StatusInternalServerError, ErrorResult(err.Error()))

	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

func RefundQueryGreen(c echo.Context) error {
	reqDto := wxpay.ReqRefundQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}

	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.RefundQuery(&reqDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

func PrePayGreen(c echo.Context) error {
	reqDto := wxpay.ReqPrepayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}

	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Prepay(&reqDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
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

	account := greenAccount{AppId: EnvParam.AppId, Key: EnvParam.Key, MchId: EnvParam.MchId,
		CertPathName: EnvParam.CertName, CertPathKey: EnvParam.CertKey, RootCa: EnvParam.RootCa,
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

func ErrorResult(errMsg string) (result model.Result) {
	result = model.Result{
		Error: model.Error{Message: errMsg},
	}
	return
}
func SuccessResult(param interface{}) (result model.Result) {
	result = model.Result{
		Success: true,
		Result:  param,
	}
	return
}
