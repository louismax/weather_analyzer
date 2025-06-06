package main

import (
	"github.com/louismax/weather_analyzer/analyzer"
	"github.com/louismax/weather_analyzer/qweather"
	"strconv"
	"testing"
)

func TestWeatherAnalyzer(t *testing.T) {
	// 创建测试数据
	conditions := []analyzer.WeatherCondition{
		{Time: "2023-10-01 00:00", Temperature: 20.0, Condition: "晴", Humidity: 60.0, WindSpeed: 5.0, Precipitation: 0.0},
		{Time: "2023-10-01 01:00", Temperature: 19.0, Condition: "多云", Humidity: 65.0, WindSpeed: 6.0, Precipitation: 0.0},
		{Time: "2023-10-01 02:00", Temperature: 18.0, Condition: "小雨", Humidity: 70.0, WindSpeed: 7.0, Precipitation: 0.5},
		{Time: "2023-10-01 03:00", Temperature: 17.0, Condition: "中雨", Humidity: 75.0, WindSpeed: 8.0, Precipitation: 2.0},
		{Time: "2023-10-01 04:00", Temperature: 16.0, Condition: "大雨", Humidity: 80.0, WindSpeed: 9.0, Precipitation: 5.0},
		{Time: "2023-10-01 05:00", Temperature: 15.0, Condition: "雷阵雨", Humidity: 85.0, WindSpeed: 10.0, Precipitation: 8.0},
		{Time: "2023-10-01 06:00", Temperature: 14.0, Condition: "暴雨", Humidity: 90.0, WindSpeed: 11.0, Precipitation: 12.0},
		{Time: "2023-10-01 07:00", Temperature: 13.0, Condition: "大暴雨", Humidity: 95.0, WindSpeed: 12.0, Precipitation: 15.0},
		{Time: "2023-10-01 08:00", Temperature: 12.0, Condition: "特大暴雨", Humidity: 100.0, WindSpeed: 13.0, Precipitation: 20.0},
		{Time: "2023-10-01 09:00", Temperature: 11.0, Condition: "强雷阵雨", Humidity: 95.0, WindSpeed: 14.0, Precipitation: 18.0},
		{Time: "2023-10-01 10:00", Temperature: 10.0, Condition: "雷阵雨伴有冰雹", Humidity: 90.0, WindSpeed: 15.0, Precipitation: 16.0},
		{Time: "2023-10-01 11:00", Temperature: 9.0, Condition: "极端降雨", Humidity: 85.0, WindSpeed: 16.0, Precipitation: 14.0},
		{Time: "2023-10-01 12:00", Temperature: 8.0, Condition: "强沙尘暴", Humidity: 80.0, WindSpeed: 17.0, Precipitation: 0.0},
		{Time: "2023-10-01 13:00", Temperature: 7.0, Condition: "特强浓雾", Humidity: 75.0, WindSpeed: 18.0, Precipitation: 0.0},
		{Time: "2023-10-01 14:00", Temperature: 6.0, Condition: "严重霾", Humidity: 70.0, WindSpeed: 19.0, Precipitation: 0.0},
		{Time: "2023-10-01 15:00", Temperature: 5.0, Condition: "冻雨", Humidity: 65.0, WindSpeed: 20.0, Precipitation: 1.0},
		{Time: "2023-10-01 16:00", Temperature: 4.0, Condition: "大雪", Humidity: 60.0, WindSpeed: 21.0, Precipitation: 3.0},
		{Time: "2023-10-01 17:00", Temperature: 3.0, Condition: "暴雪", Humidity: 55.0, WindSpeed: 22.0, Precipitation: 6.0},
		{Time: "2023-10-01 18:00", Temperature: 2.0, Condition: "大到暴雪", Humidity: 50.0, WindSpeed: 23.0, Precipitation: 9.0},
		{Time: "2023-10-01 19:00", Temperature: 1.0, Condition: "中到大雪", Humidity: 45.0, WindSpeed: 24.0, Precipitation: 7.0},
		{Time: "2023-10-01 20:00", Temperature: 0.0, Condition: "小到中雪", Humidity: 40.0, WindSpeed: 25.0, Precipitation: 4.0},
		{Time: "2023-10-01 21:00", Temperature: -1.0, Condition: "阵雪", Humidity: 35.0, WindSpeed: 26.0, Precipitation: 2.0},
		{Time: "2023-10-01 22:00", Temperature: -2.0, Condition: "雪", Humidity: 30.0, WindSpeed: 27.0, Precipitation: 1.0},
		{Time: "2023-10-01 23:00", Temperature: -3.0, Condition: "雨夹雪", Humidity: 25.0, WindSpeed: 28.0, Precipitation: 0.5},
	}

	// 自定义天气状况权重,一般不建议调整
	// customWeights := map[string]float64{
	// 	"特大暴雨": 1.0,
	// 	"大暴雨":  0.9,
	// 	"暴雨":   0.8,
	// 	"大雨":   0.7,
	// 	"中雨":   0.6,
	// 	"小雨":   0.5,
	// 	"晴":    0.1,
	// 	"多云":   0.2,
	// }

	// 创建天气分析器
	wa, err := analyzer.NewWeatherAnalyzer(conditions)
	if err != nil {
		t.Fatalf("创建天气分析器失败: %v", err)
	}

	// 执行分析
	result, err := wa.Analyze()
	if err != nil {
		t.Fatalf("分析天气数据失败: %v", err)
	}

	// 打印结果
	t.Logf("天气分析结果:")
	t.Logf("主导天气状况: %s", result.DominantCondition)
	t.Logf("平均温度: %.2f°C", result.AverageTemperature)
	t.Logf("总降水量: %.2f mm", result.TotalPrecipitation)
	t.Logf("降水持续小时数: %d", result.PrecipitationHours)
	t.Logf("平均风速: %.2f m/s", result.AverageWindSpeed)
	t.Logf("最大风速: %.2f m/s", result.MaxWindSpeed)
	t.Logf("天气描述: %s", result.Description)
}

func TestQWeatherHistoryAnalyzer(t *testing.T) {
	client, err := qweather.NewQWeatherApiClient("YOUR_KEY_ID", "YOUR_PROJECT_ID", "YOUR_API_HOST", "./privateKey.pem")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Request(qweather.APIHistoricalWeather, map[string]string{
		"location": "101250109",
		"date":     "20250604",
	})
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.HistoricalWeatherResult()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)

	// 创建测试数据
	conditions := make([]analyzer.WeatherCondition, 0)
	for _, v := range res.WeatherHourly {
		f, _ := strconv.ParseFloat(v.Temp, 64)
		humidity, _ := strconv.ParseFloat(v.Humidity, 64)
		windSpeed, _ := strconv.ParseFloat(v.WindSpeed, 64)
		prec_ip, _ := strconv.ParseFloat(v.Precip, 64)
		conditions = append(conditions, analyzer.WeatherCondition{
			Time:          v.Time,
			Temperature:   f,
			Condition:     v.Text,
			Humidity:      humidity,
			WindSpeed:     windSpeed,
			Precipitation: prec_ip,
		})
	}

	// 创建天气分析器
	wa, err := analyzer.NewWeatherAnalyzer(conditions)
	if err != nil {
		t.Fatalf("创建天气分析器失败: %v", err)
	}

	// 执行分析
	result, err := wa.Analyze()
	if err != nil {
		t.Fatalf("分析天气数据失败: %v", err)
	}

	// 打印结果
	t.Logf("天气分析结果:")
	t.Logf("主导天气状况: %s", result.DominantCondition)
	t.Logf("平均温度: %.2f°C", result.AverageTemperature)
	t.Logf("总降水量: %.2f mm", result.TotalPrecipitation)
	t.Logf("降水持续小时数: %d", result.PrecipitationHours)
	t.Logf("平均风速: %.2f m/s", result.AverageWindSpeed)
	t.Logf("最大风速: %.2f m/s", result.MaxWindSpeed)
	t.Logf("天气描述: %s", result.Description)
}
