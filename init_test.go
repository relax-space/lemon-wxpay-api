package wxpayapi

import (
	"flag"
	"os"
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

func init() {
	flag.Parse()
	EnvParam = &EnvParamDto{
		AppEnv:   *appEnv,
		AppId:    *appId,
		Key:      *key,
		MchId:    *mchId,
		CertName: *certName,
		CertKey:  *certKey,
		RootCa:   *rootCa,
	}
}
