package qweather

import (
	"crypto/ed25519"
	"testing"
)

func TestNewQWeatherApiClient(t *testing.T) {
	client, err := NewQWeatherApiClient("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", "./privateKey.pem")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client.Token)
}

func TestNewQWeatherApiClientByPKString(t *testing.T) {
	pkStr := "-----BEGIN PRIVATE KEY-----\nYOUR_PRIVATE_KEY\n-----END PRIVATE KEY-----"
	client, err := NewQWeatherApiClientByPKString("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", pkStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client.Token)
}

func TestNewQWeatherApiClientByPKED(t *testing.T) {
	pkED := ed25519.PrivateKey{} //YOUR_PRIVATE_KEY
	client, err := NewQWeatherApiClientByPKED("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", pkED)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client.Token)
}

func TestQWeatherApiClientRequest(t *testing.T) {
	client, err := NewQWeatherApiClient("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", "./privateKey.pem")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Request(APIGeoCityLookup, map[string]string{
		"location": "岳麓",
		"adm":      "湖南",
		"range":    "cn",
		"lang":     "zh",
	})
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.GeoCityLookupResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

func TestWeatherIconCode(t *testing.T) {
	c := ApiClient{}
	t.Log(c.GetWeatherIconCode()["晴"])
}
