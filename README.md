# weather_analyzer

[简体中文](README.md)
### weather_analyzer是一个Golang编写的天气状况分析工具, 同时也可作为Golang版本的和风天气API SDK；
- ### analyzer： 天气变化瞬息万变，天气数据也是如此； 此项目用于天气状况分析，例如根据某天的逐小时天气数据，分析当天整体天气情况！
- ### qweather：和风天气API Golang SDK，方便开发者快速接入和风天气API，实现天气数据获取、天气预警推送等功能。

[![Go Report Card](https://goreportcard.com/badge/github.com/louismax/weather_analyzer)](https://goreportcard.com/report/github.com/louismax/weather_analyzer)
[![GoDoc](https://godoc.org/github.com/louismax/weather_analyzer?status.svg)](https://godoc.org/github.com/louismax/weather_analyzer)
[![GitHub release](https://img.shields.io/github/tag/louismax/weather_analyzer.svg)](https://github.com/louismax/weather_analyzer/releases)
[![GitHub license](https://img.shields.io/github/license/louismax/weather_analyzer.svg)](https://github.com/louismax/weather_analyzer/blob/master/LICENSE)
[![GitHub Repo Size](https://img.shields.io/github/repo-size/louismax/weather_analyzer.svg)](https://img.shields.io/github/repo-size/louismax/weather_analyzer.svg)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/louismax/weather_analyzer.svg)](https://img.shields.io/github/last-commit/louismax/weather_analyzer.svg)

## 安装
`go get github.com/louismax/weather_analyzer`

## 🚀 analyzer使用
如果你有某一天的逐小时天气数据，希望科学的分析得到当天整体的天气状况，可以使用本项目的天气分析器获取；例如和风天气的时光机天气(历史天气)API,返回某一天的天气状况中是没有当天的整体的天气状况描述的，由于天气数据瞬息万变的原因，只有逐小时天气中又天气状况描述，而实际项目中又不方便显示逐小时天气状况，故而使用逐小时天气数据来分析总结当天的天气状况
### 创建天气分析器
需要传入天气数据集合，一般为时间段内的逐小时天气数据
```go
wa, err := analyzer.NewWeatherAnalyzer([]analyzer.WeatherCondition)
if err != nil {
    t.Fatalf("创建天气分析器失败: %v", err)
}
```
### 设置自定义天气状况权重
天气状况分析的核心逻辑，除了根据某种天气现象的出现情况，还应还为天气状况定义权重才更加科学，权重设置原则：
1. 极端天气（如暴雨、强雷阵雨等）权重最高
2. 伴有特殊现象（如冰雹、雷电）的天气权重较高
3. 降水强度大的天气权重较高
4. 持续时间长的天气权重较高

天气分析器中已经定义好了一套默认天气权重，具体可查看[👉默认配置](DefaultCfg.md)，如果默认配置无法满足或认为默认权重不科学，也可以根据实际需求自定义天气权重
```go
wa.SetCustomWeights(map[string]float64{
    "晴":  0.4,
    "多云": 0.6,
    "阴":  0.7,
})
```
### 设置自定义降水量阈值
对于部分天气状况,还需要根据降水量调整权重

天气分析器中已经定义好了一套默认降水量阈值，具体可查看[👉默认配置](DefaultCfg.md)，如果默认配置无法满足或认为默认配置不科学，也可以根据实际需求自定义降水量阈值
```go
wa.SetCustomPrecipitationThresholds(map[string]float64{
    "特大暴雨": 120.0,
    "大暴雨":  90.0,
    "暴雨":   60.0,
})
```

### 设置自定义风速阈值
对于部分天气状况,还需要根据风速调整权重

天气分析器中已经定义好了一套默认风速阈值，具体可查看[👉默认配置](DefaultCfg.md)，如果默认配置无法满足或认为默认配置不科学，也可以根据实际需求自定义风速阈值
```go
wa.SetCustomWindSpeedThresholds(map[string]float64{
    "强沙尘暴": 25.0,
    "沙尘暴":  20.0,
    "扬沙":   15.0,
})
```

### 获取分析结果
```go
result, err := analyzer.Analyze()
if err != nil {
    t.Fatalf("分析天气状况失败: %v", err)
}
```
打印结果
```text
    local_test.go:81: 天气分析结果:
    local_test.go:82: 主导天气状况: 多云
    local_test.go:83: 平均温度: 25.92°C
    local_test.go:84: 总降水量: 0.00 mm
    local_test.go:85: 降水持续小时数: 0
    local_test.go:86: 平均风速: 6.71 m/s
    local_test.go:87: 最大风速: 16.00 m/s
    local_test.go:88: 天气描述: 今日天气以多云为主，平均温度25.9°C，平均风速6.7米/秒，最大风速16.0米/秒。期间还出现阴、晴。
```

## 🚀 qweather使用
qweather是和风天气API Golang SDK，方便开发者快速接入和风天气API，实现天气数据获取、天气预警推送等功能。
### 创建一个新的和风天气ApiClient实例
需要传入和风天气项目ID，凭证ID，API主机地址，以及私钥文件路径
```go
client, err := qweather.NewQWeatherApiClient("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", "./privateKey.pem")
if err != nil {
    t.Fatal(err)
}
```
也可通过私钥字符串或私钥创建
```go
//通过私钥字符串创建
pkStr := "-----BEGIN PRIVATE KEY-----\nYOUR_PRIVATE_KEY\n-----END PRIVATE KEY-----"
client, err := qweather.NewQWeatherApiClientByPKString("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", pkStr)
if err != nil {
    t.Fatal(err)
}

//通过私钥创建
pkED := ed25519.PrivateKey{} //YOUR_PRIVATE_KEY
client, err := qweather.NewQWeatherApiClientByPKED("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", pkED)
if err != nil {
    t.Fatal(err)
}
```

### 调用和风天气API
```go
resp, err := client.Request("/geo/v2/city/lookup", map[string]string{
    "location": "岳麓",
    "adm":      "湖南",
    "range":    "cn",
    "lang":     "zh",
})
if err != nil {
    t.Fatal(err)
}
```
需要传入和风天气API接口路径和请求参数（即路径?之后的参数）；需要注意的是，部分接口路径包含动态参数，需要自行处理，例如时光机API，接口路径中包含{days}

### 和风天气API结果解析
Request返回结果对于部分常用的API已经实现的结构体解析，可以直接使用
```go
// GEO城市查询结果
result, err := resp.GeoCityLookupResult()
// GEO热门城市查询结果
result, err := resp.GeoTopCityResult()
// POI查询结果(含poi range)
result, err := resp.GeoPoiResult()
// 实时天气查询结果
result, err := resp.NowWeatherResult()
// 每日预报天气查询结果
result, err := resp.ForecastWeatherResult()
//  时光机天气(历史天气)查询结果
result, err := resp.HistoricalWeatherResult()
```
其他未封装的请求结果结构体,可自行定义,并对resp.Body进行JSON解析
```go
//根据API文档自行定义的结构体
customResult := struct {
	A string `json:"a"`
	B  Float64 `json:"b"`
}{}
if err := json.Unmarshal(resp.Body, &customResult); err != nil {
    return nil, err
}
```

## 参考资料
* [中国气象-天气分析的内容和方法](http://stream1.cmatc.cn/cmatcvod/12/tqx/first_points.html)
* [和风天气开发服务](https://dev.qweather.com/docs/api/)

## 协议
Apache-License2.0。有关更多信息，请参见[协议文件](LICENSE)。
