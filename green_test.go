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
