package analyzer

import (
	"testing"
)

func TestWeatherAnalyzer(t *testing.T) {
	// 创建测试数据
	conditions := []WeatherCondition{
		{Time: "2024-01-01 00:00", Temperature: 25.0, Condition: "晴", Humidity: 60.0, WindSpeed: 5.0, Precipitation: 0.0},
		{Time: "2024-01-01 01:00", Temperature: 24.0, Condition: "多云", Humidity: 65.0, WindSpeed: 6.0, Precipitation: 0.0},
		{Time: "2024-01-01 02:00", Temperature: 23.0, Condition: "阴", Humidity: 70.0, WindSpeed: 7.0, Precipitation: 0.0},
		{Time: "2024-01-01 03:00", Temperature: 22.0, Condition: "小雨", Humidity: 75.0, WindSpeed: 8.0, Precipitation: 0.1},
		{Time: "2024-01-01 04:00", Temperature: 21.0, Condition: "中雨", Humidity: 80.0, WindSpeed: 9.0, Precipitation: 10.0},
		{Time: "2024-01-01 05:00", Temperature: 20.0, Condition: "大雨", Humidity: 85.0, WindSpeed: 10.0, Precipitation: 25.0},
		{Time: "2024-01-01 06:00", Temperature: 19.0, Condition: "暴雨", Humidity: 90.0, WindSpeed: 11.0, Precipitation: 50.0},
		{Time: "2024-01-01 07:00", Temperature: 18.0, Condition: "大暴雨", Humidity: 95.0, WindSpeed: 12.0, Precipitation: 70.0},
		{Time: "2024-01-01 08:00", Temperature: 17.0, Condition: "特大暴雨", Humidity: 100.0, WindSpeed: 13.0, Precipitation: 100.0},
		{Time: "2024-01-01 09:00", Temperature: 16.0, Condition: "雷阵雨", Humidity: 95.0, WindSpeed: 14.0, Precipitation: 30.0},
		{Time: "2024-01-01 10:00", Temperature: 15.0, Condition: "强雷阵雨", Humidity: 90.0, WindSpeed: 15.0, Precipitation: 40.0},
		{Time: "2024-01-01 11:00", Temperature: 14.0, Condition: "雷阵雨伴有冰雹", Humidity: 85.0, WindSpeed: 16.0, Precipitation: 20.0},
		{Time: "2024-01-01 12:00", Temperature: 13.0, Condition: "极端降雨", Humidity: 80.0, WindSpeed: 17.0, Precipitation: 80.0},
		{Time: "2024-01-01 13:00", Temperature: 12.0, Condition: "强阵雨", Humidity: 75.0, WindSpeed: 18.0, Precipitation: 15.0},
		{Time: "2024-01-01 14:00", Temperature: 11.0, Condition: "暴雪", Humidity: 70.0, WindSpeed: 19.0, Precipitation: 60.0},
		{Time: "2024-01-01 15:00", Temperature: 10.0, Condition: "强沙尘暴", Humidity: 65.0, WindSpeed: 20.8, Precipitation: 0.0},
		{Time: "2024-01-01 16:00", Temperature: 9.0, Condition: "沙尘暴", Humidity: 60.0, WindSpeed: 17.2, Precipitation: 0.0},
		{Time: "2024-01-01 17:00", Temperature: 8.0, Condition: "扬沙", Humidity: 55.0, WindSpeed: 10.8, Precipitation: 0.0},
		{Time: "2024-01-01 18:00", Temperature: 7.0, Condition: "浮尘", Humidity: 50.0, WindSpeed: 5.5, Precipitation: 0.0},
		{Time: "2024-01-01 19:00", Temperature: 6.0, Condition: "特强浓雾", Humidity: 45.0, WindSpeed: 4.0, Precipitation: 0.0},
		{Time: "2024-01-01 20:00", Temperature: 5.0, Condition: "强浓雾", Humidity: 40.0, WindSpeed: 3.0, Precipitation: 0.0},
		{Time: "2024-01-01 21:00", Temperature: 4.0, Condition: "严重霾", Humidity: 35.0, WindSpeed: 2.0, Precipitation: 0.0},
		{Time: "2024-01-01 22:00", Temperature: 3.0, Condition: "重度霾", Humidity: 30.0, WindSpeed: 1.0, Precipitation: 0.0},
		{Time: "2024-01-01 23:00", Temperature: 2.0, Condition: "中度霾", Humidity: 25.0, WindSpeed: 0.0, Precipitation: 0.0},
	}

	// 创建天气分析器
	analyzer, err := NewWeatherAnalyzer(conditions)
	if err != nil {
		t.Fatalf("创建天气分析器失败: %v", err)
	}

	// 设置自定义权重
	customWeights := map[string]float64{
		"晴":  0.4,
		"多云": 0.6,
		"阴":  0.7,
	}
	analyzer.SetCustomWeights(customWeights)

	// 设置自定义降水量阈值
	customPrecipitationThresholds := map[string]float64{
		"特大暴雨": 120.0,
		"大暴雨":  90.0,
		"暴雨":   60.0,
	}
	analyzer.SetCustomPrecipitationThresholds(customPrecipitationThresholds)

	// 设置自定义风速阈值
	customWindSpeedThresholds := map[string]float64{
		"强沙尘暴": 25.0,
		"沙尘暴":  20.0,
		"扬沙":   15.0,
	}
	analyzer.SetCustomWindSpeedThresholds(customWindSpeedThresholds)

	// 分析天气状况
	result, err := analyzer.Analyze()
	if err != nil {
		t.Fatalf("分析天气状况失败: %v", err)
	}

	// 验证分析结果
	expected := &WeatherAnalysisResult{
		DominantCondition:  "特大暴雨",
		OtherConditions:    []string{"大暴雨", "极端降雨", "强沙尘暴"},
		AverageTemperature: 13.5,
		TotalPrecipitation: 500.2,
		MaxPrecipitation:   100.0,
		PrecipitationHours: 12,
		AverageWindSpeed:   10.0,
		MaxWindSpeed:       20.8,
		ConditionWeights: map[string]float64{
			"特大暴雨": 1.0,
			"大暴雨":  0.995,
			"极端降雨": 0.99,
			"强沙尘暴": 0.93,
		},
		Description: "今日天气以特大暴雨为主，平均温度13.5°C，总降水量500.2毫米，降水持续12小时，最大小时降水量100.0毫米，平均风速10.0米/秒，最大风速20.8米/秒。期间还出现大暴雨、极端降雨、强沙尘暴。",
	}

	if result.DominantCondition != expected.DominantCondition {
		t.Errorf("主要天气状况错误，期望 %s，实际 %s", expected.DominantCondition, result.DominantCondition)
	}

	if len(result.OtherConditions) != len(expected.OtherConditions) {
		t.Errorf("其他重要天气状况数量错误，期望 %d，实际 %d", len(expected.OtherConditions), len(result.OtherConditions))
	}

	if result.AverageTemperature != expected.AverageTemperature {
		t.Errorf("平均温度错误，期望 %.1f，实际 %.1f", expected.AverageTemperature, result.AverageTemperature)
	}

	if result.TotalPrecipitation != expected.TotalPrecipitation {
		t.Errorf("总降水量错误，期望 %.1f，实际 %.1f", expected.TotalPrecipitation, result.TotalPrecipitation)
	}

	if result.MaxPrecipitation != expected.MaxPrecipitation {
		t.Errorf("最大小时降水量错误，期望 %.1f，实际 %.1f", expected.MaxPrecipitation, result.MaxPrecipitation)
	}

	if result.PrecipitationHours != expected.PrecipitationHours {
		t.Errorf("降水持续小时数错误，期望 %d，实际 %d", expected.PrecipitationHours, result.PrecipitationHours)
	}

	if result.AverageWindSpeed != expected.AverageWindSpeed {
		t.Errorf("平均风速错误，期望 %.1f，实际 %.1f", expected.AverageWindSpeed, result.AverageWindSpeed)
	}

	if result.MaxWindSpeed != expected.MaxWindSpeed {
		t.Errorf("最大风速错误，期望 %.1f，实际 %.1f", expected.MaxWindSpeed, result.MaxWindSpeed)
	}

	if len(result.ConditionWeights) != len(expected.ConditionWeights) {
		t.Errorf("天气状况权重数量错误，期望 %d，实际 %d", len(expected.ConditionWeights), len(result.ConditionWeights))
	}

}
