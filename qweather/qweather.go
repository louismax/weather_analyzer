package qweather

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/louismax/weather_analyzer/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type ApiClient struct {
	SHeader struct {
		Alg string `json:"alg"`
		Kid string `json:"kid"`
	}
	SPayload struct {
		Sub string `json:"sub"`
		Iat int64  `json:"iat"`
		Exp int64  `json:"exp"`
	}
	PrivateKey ed25519.PrivateKey
	ApiHost    string
	Token      string
}

// NewQWeatherApiClient 创建一个新的和风天气ApiClient实例
func NewQWeatherApiClient(kId, subId, apiHost, PrivateKeyPath string) (*ApiClient, error) {
	//读取私钥文件
	privateKeyPEM, err := os.ReadFile(PrivateKeyPath)
	if err != nil {
		utils.PrintErrorLog(err.Error())
		return nil, &utils.WeatherError{
			Code:    utils.ErrReadFile,
			Message: fmt.Sprintf("读取私钥文件失败.%s", err.Error()),
		}
	}
	//解析私钥
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, &utils.WeatherError{
			Code:    utils.ErrPrivateKeyInvalid,
			Message: "私钥无效",
		}
	}
	//PKCS#8解析
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		utils.PrintErrorLog(err.Error())
		return nil, &utils.WeatherError{
			Code:    utils.ErrInvalidInput,
			Message: fmt.Sprintf("PKCS#8解析失败.%s", err.Error()),
		}
	}
	ed25519Key, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, &utils.WeatherError{
			Code:    utils.ErrInvalidInput,
			Message: "Not an ED25519 private key",
		}
	}
	return initQWeatherApiClient(kId, subId, apiHost, ed25519Key)
}

// NewQWeatherApiClientByPKString 创建一个新的和风天气ApiClient实例(通过PrivateKey明文字符串)
func NewQWeatherApiClientByPKString(kId, subId, apiHost, PrivateKey string) (*ApiClient, error) {
	//解析私钥
	block, _ := pem.Decode([]byte(PrivateKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, &utils.WeatherError{
			Code:    utils.ErrPrivateKeyInvalid,
			Message: "私钥无效",
		}
	}
	//PKCS#8解析
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		utils.PrintErrorLog(err.Error())
		return nil, &utils.WeatherError{
			Code:    utils.ErrInvalidInput,
			Message: fmt.Sprintf("PKCS#8解析失败.%s", err.Error()),
		}
	}
	ed25519Key, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, &utils.WeatherError{
			Code:    utils.ErrInvalidInput,
			Message: "Not an ED25519 private key",
		}
	}
	return initQWeatherApiClient(kId, subId, apiHost, ed25519Key)
}

// NewQWeatherApiClientByPKED 创建一个新的和风天气ApiClient实例(通过PrivateKey ed25519.PrivateKey)
func NewQWeatherApiClientByPKED(kId, subId, apiHost string, PrivateKey ed25519.PrivateKey) (*ApiClient, error) {
	return initQWeatherApiClient(kId, subId, apiHost, PrivateKey)
}

func initQWeatherApiClient(kId, subId, apiHost string, PrivateKey ed25519.PrivateKey) (*ApiClient, error) {
	cli := ApiClient{
		SHeader: struct {
			Alg string `json:"alg"`
			Kid string `json:"kid"`
		}{
			Alg: "EdDSA",
			Kid: kId,
		},
		SPayload: struct {
			Sub string `json:"sub"`
			Iat int64  `json:"iat"`
			Exp int64  `json:"exp"`
		}{
			Sub: subId,
		},
		ApiHost:    apiHost,
		PrivateKey: PrivateKey,
	}
	cli.sign()
	return &cli, nil
}

func (c *ApiClient) Request(methodPath string, params map[string]string) (*ResultQWeather, error) {
	_url := fmt.Sprintf("https://%s%s", c.ApiHost, methodPath)
	if len(params) > 0 {
		_url += "?"
		inx := 1
		for k, v := range params {
			if inx != len(params) {
				_url += fmt.Sprintf("%s=%s&", k, url.QueryEscape(v))
			} else {
				_url += fmt.Sprintf("%s=%s", k, url.QueryEscape(v))
			}
			inx++
		}
	}
	if c.SPayload.Exp >= time.Now().Unix() {
		c.sign()
	}
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, _url, nil)
	if err != nil {
		utils.PrintErrorLog("请求创建失败,error:%+v", err)
		return nil, &utils.WeatherError{
			Code:    utils.ErrRequestFailed,
			Message: fmt.Sprintf("请求创建失败,error:%+v", err),
		}
	}
	request.Header.Add("Authorization", c.Token)
	response, err := client.Do(request)
	if err != nil {
		utils.PrintErrorLog("请求发送失败,error:%+v", err)
		return nil, &utils.WeatherError{
			Code:    utils.ErrRequestFailed,
			Message: fmt.Sprintf("请求发送失败,error:%+v", err),
		}
	}
	defer func() {
		_ = response.Body.Close()
	}()
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		utils.PrintErrorLog("请求结果解析失败,error:%+v", err)
		return nil, &utils.WeatherError{
			Code:    utils.ErrRequestFailed,
			Message: fmt.Sprintf("请求结果解析失败,error:%+v", err),
		}
	}
	return &ResultQWeather{
		Body: resp,
	}, nil
}

// sign 签名
func (c *ApiClient) sign() {
	c.SPayload.Iat = time.Now().Add(time.Minute * -1).Unix()
	c.SPayload.Exp = time.Now().Add(time.Minute * 30).Unix()
	HeaderBase64URL := utils.Struct2Base64URL(c.SHeader)
	PayloadBase64URL := utils.Struct2Base64URL(c.SPayload)
	//数据加密
	sig := ed25519.Sign(c.PrivateKey, []byte(HeaderBase64URL+"."+PayloadBase64URL))
	//数据Base64编码
	SignatureBase64URL := base64.URLEncoding.EncodeToString(sig)

	c.Token = fmt.Sprintf("Bearer %s.%s.%s", HeaderBase64URL, PayloadBase64URL, SignatureBase64URL)
}
