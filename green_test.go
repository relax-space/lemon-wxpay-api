package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/relax-space/go-kit/test"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/model"
)

func Test_PayGreen(t *testing.T) {
	bodyStr := `
	{
		"auth_code":"135298324463700425",
		"body":"xiaoxinmiao test",
		"total_fee":1
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/pay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, PayGreen(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_RefundGreen(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no":"147688874645492354650",
		"refund_fee":1
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/refund", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, RefundGreen(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_ReverseGreen(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no":"143420620288156126697"
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/reverse", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, ReverseGreen(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_QueryGreen(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no":"14201711085205823413229775520"
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/query", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, QueryGreen(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_RefundQueryGreen(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no":"144650782494807835413"
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/refundquery", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, RefundQueryGreen(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_PrePayGreen(t *testing.T) {
	bodyStr := `
	{
		"body":"xiaomiao test",
		"total_fee":1,
		"trade_type":"JSAPI",
		"notify_url":"http://xiao.xinmiao.com",
		"openid":"os2u9uPKLkCKL08FwCM6hQAQ_LtI"
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/prepay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, PrePayGreen(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_NotifyGreen(t *testing.T) {
	xmlBody := `<xml>
	<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
	<attach><![CDATA[{"sub_notify_url":"https://baidu.com"}]]></attach>
	<bank_type><![CDATA[CFT]]></bank_type>
	<fee_type><![CDATA[CNY]]></fee_type>
	<is_subscribe><![CDATA[Y]]></is_subscribe>
	<mch_id><![CDATA[10000100]]></mch_id>
	<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
	<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
	<out_trade_no><![CDATA[1409811653]]></out_trade_no>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
	<sub_mch_id><![CDATA[10000100]]></sub_mch_id>
	<time_end><![CDATA[20140903131540]]></time_end>
	<total_fee>1</total_fee>
	<trade_type><![CDATA[JSAPI]]></trade_type>
	<transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
 </xml>`
	req, err := http.NewRequest(echo.POST, "/v1/green/notify", strings.NewReader(xmlBody))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, NotifyGreen(c))
	test.Equals(t, http.StatusOK, rec.Code)
	fmt.Printf("%+v", string(rec.Body.Bytes()))

}
