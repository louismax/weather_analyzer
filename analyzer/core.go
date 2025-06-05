package analyzer

import "fmt"

// WeatherCondition 表示天气状况
type WeatherCondition struct {
	// Time 表示观测时间
	Time string
	// Temperature 表示温度，单位为摄氏度
	Temperature float64
	// Condition 表示天气状况，如：晴、多云、雨等
	Condition string
	// Humidity 表示相对湿度，范围为0-100
	Humidity float64
	// WindSpeed 表示风速，单位为米/秒
	WindSpeed float64
	// Precipitation 表示降水量，单位为毫米
	Precipitation float64
}

// WeatherAnalyzer 天气分析器
type WeatherAnalyzer struct {
	conditions []WeatherCondition
	// 天气状况权重映射
	conditionWeights map[string]float64
	// 降水量阈值（毫米/小时）
	precipitationThresholds map[string]float64
	// 风速阈值（米/秒）
	windSpeedThresholds map[string]float64
}

// WeatherAnalysisResult 天气分析结果
type WeatherAnalysisResult struct {
	// 主要天气状况
	DominantCondition string
	// 其他重要天气状况
	OtherConditions []string
	// 平均温度（摄氏度）
	AverageTemperature float64
	// 总降水量（毫米）
	TotalPrecipitation float64
	// 最大小时降水量（毫米）
	MaxPrecipitation float64
	// 降水持续小时数
	PrecipitationHours int
	// 平均风速（米/秒）
	AverageWindSpeed float64
	// 最大风速（米/秒）
	MaxWindSpeed float64
	// 天气状况权重统计
	ConditionWeights map[string]float64
	// 天气描述文本
	Description string
}

// NewWeatherAnalyzer 创建新的天气分析器
func NewWeatherAnalyzer(conditions []WeatherCondition) (*WeatherAnalyzer, error) {
	if conditions == nil {
		return nil, &WeatherError{
			Code:    ErrInvalidInput,
			Message: "天气数据不能为空",
		}
	}

	if len(conditions) == 0 {
		return nil, &WeatherError{
			Code:    ErrEmptyData,
			Message: "天气数据列表为空",
		}
	}

	// 验证每个天气条件
	for i, condition := range conditions {
		if err := validateWeatherCondition(condition); err != nil {
			return nil, &WeatherError{
				Code:    ErrInvalidInput,
				Message: fmt.Sprintf("第%d个天气数据无效", i+1),
				Err:     err,
			}
		}
	}

	// 初始化天气状况权重
	// 权重设置原则：
	// 1. 极端天气（如暴雨、强雷阵雨等）权重最高
	// 2. 伴有特殊现象（如冰雹、雷电）的天气权重较高
	// 3. 降水强度大的天气权重较高
	// 4. 持续时间长的天气权重较高
	weights := map[string]float64{
		// 极端天气（权重 0.95-1.00）
		"特大暴雨":    1.00,  // 最强降水，持续时间长，影响范围大
		"大暴雨":     0.995, // 极强降水，持续时间长
		"极端降雨":    0.99,  // 极端降水，影响范围大
		"雷阵雨伴有冰雹": 0.98,  // 伴有冰雹，危险性极高
		"强雷阵雨":    0.97,  // 强降水+雷电，危险性高
		"强阵雨":     0.96,  // 强降水，持续时间短
		"暴雨":      0.955, // 强降水，持续时间长
		"暴雪":      0.95,  // 强降雪，持续时间长，影响大
		"强沙尘暴":    0.93,  // 极强沙尘，能见度极低
		"特强浓雾":    0.92,  // 能见度极低，影响极大

		// 较强天气（权重 0.85-0.94）
		"雷阵雨": 0.92, // 降水+雷电，危险性较高
		"大雨":  0.90, // 较强降水，持续时间长
		"中雨":  0.88, // 中等降水，持续时间长
		"冻雨":  0.87, // 特殊降水，易造成道路结冰
		"大雪":  0.86, // 较强降雪，持续时间长
		"沙尘暴": 0.85, // 强沙尘，能见度低
		"强浓雾": 0.83, // 能见度很低
		"严重霾": 0.81, // 空气质量极差

		// 过渡性天气（权重 0.80-0.89）
		"大暴雨到特大暴雨": 0.89, // 极强降水过渡
		"暴雨到大暴雨":   0.87, // 强降水过渡
		"大到暴雨":     0.85, // 较强降水过渡
		"中到大雨":     0.83, // 中等降水过渡
		"小到中雨":     0.81, // 弱降水过渡
		"大到暴雪":     0.84, // 较强降雪过渡
		"中到大雪":     0.82, // 中等降雪过渡
		"小到中雪":     0.80, // 弱降雪过渡

		// 一般天气（权重 0.70-0.84）
		"阵雨":     0.82, // 短时降水
		"雨":      0.80, // 基础降水天气
		"小雨":     0.78, // 弱降水，持续时间长
		"毛毛雨/细雨": 0.75, // 极弱降水，持续时间长
		"中雪":     0.77, // 中等降雪
		"小雪":     0.76, // 弱降雪
		"雪":      0.75, // 基础降雪天气
		"雨夹雪":    0.74, // 雨雪混合
		"雨雪天气":   0.73, // 雨雪交替
		"阵雨夹雪":   0.72, // 短时雨雪混合
		"阵雪":     0.71, // 短时降雪
		"重度霾":    0.79, // 空气质量很差
		"中度霾":    0.78, // 空气质量较差
		"大雾":     0.77, // 能见度低
		"浓雾":     0.76, // 能见度较低
		"雾":      0.75, // 基础雾天气
		"薄雾":     0.74, // 轻微雾
		"扬沙":     0.73, // 沙尘天气
		"浮尘":     0.72, // 轻微沙尘
		"霾":      0.71, // 基础霾天气

		// 无降水天气（权重 0.30-0.69）
		"阴":    0.65, // 云量多，影响光照
		"多云":   0.55, // 云量中等
		"晴间多云": 0.45, // 以晴为主，少量云
		"少云":   0.40, // 云量少
		"晴":    0.35, // 无云或少云
		"热":    0.50, // 高温天气
		"冷":    0.45, // 低温天气
	}

	// 初始化降水量阈值
	precipitationThresholds := map[string]float64{
		"特大暴雨": 100.0, // 24小时降水量≥100mm
		"大暴雨":  70.0,  // 24小时降水量≥70mm
		"暴雨":   50.0,  // 24小时降水量≥50mm
		"大雨":   25.0,  // 24小时降水量≥25mm
		"中雨":   10.0,  // 24小时降水量≥10mm
		"小雨":   0.1,   // 24小时降水量≥0.1mm
		"毛毛雨":  0.0,   // 24小时降水量>0mm
	}

	// 初始化风速阈值
	windSpeedThresholds := map[string]float64{
		"强沙尘暴": 20.8, // 风速≥20.8m/s（8级风）
		"沙尘暴":  17.2, // 风速≥17.2m/s（7级风）
		"扬沙":   10.8, // 风速≥10.8m/s（5级风）
		"浮尘":   5.5,  // 风速≥5.5m/s（3级风）
	}

	return &WeatherAnalyzer{
		conditions:              conditions,
		conditionWeights:        weights,
		precipitationThresholds: precipitationThresholds,
		windSpeedThresholds:     windSpeedThresholds,
	}, nil
}

// SetCustomWeights 设置自定义权重
func (wa *WeatherAnalyzer) SetCustomWeights(customWeights map[string]float64) {
	// 如果传入了自定义权重，则覆盖默认权重
	if customWeights != nil {
		for condition, weight := range customWeights {
			if _, exists := wa.conditionWeights[condition]; exists {
				fmt.Printf("\033[33m警告: 覆盖默认天气状况权重 '%s' (原值: %.3f, 新值: %.3f)\033[0m\n", condition, wa.conditionWeights[condition], weight)
			} else {
				fmt.Printf("\033[34m提示: 新增天气状况权重 '%s' (值: %.3f)\033[0m\n", condition, weight)
			}
			wa.conditionWeights[condition] = weight
		}
	}
}

// SetCustomPrecipitationThresholds 设置自定义降水量阈值
func (wa *WeatherAnalyzer) SetCustomPrecipitationThresholds(customPrecipitationThresholds map[string]float64) {
	// 如果传入了自定义降水量阈值，则覆盖默认阈值
	if customPrecipitationThresholds != nil {
		for condition, threshold := range customPrecipitationThresholds {
			if _, exists := wa.precipitationThresholds[condition]; exists {
				fmt.Printf("\033[33m警告: 覆盖默认降水量阈值 '%s' (原值: %.1f, 新值: %.1f)\033[0m\n", condition, wa.precipitationThresholds[condition], threshold)
			} else {
				fmt.Printf("\033[34m提示: 新增降水量阈值 '%s' (值: %.1f)\033[0m\n", condition, threshold)
			}
			wa.precipitationThresholds[condition] = threshold
		}
	}
}

// SetCustomWindSpeedThresholds 设置自定义风速阈值
func (wa *WeatherAnalyzer) SetCustomWindSpeedThresholds(customWindSpeedThresholds map[string]float64) {
	// 如果传入了自定义风速阈值，则覆盖默认阈值
	if customWindSpeedThresholds != nil {
		for condition, threshold := range customWindSpeedThresholds {
			if _, exists := wa.windSpeedThresholds[condition]; exists {
				fmt.Printf("\033[33m警告: 覆盖默认风速阈值 '%s' (原值: %.1f, 新值: %.1f)\033[0m\n", condition, wa.windSpeedThresholds[condition], threshold)
			} else {
				fmt.Printf("\033[34m提示: 新增风速阈值 '%s' (值: %.1f)\033[0m\n", condition, threshold)
			}
			wa.windSpeedThresholds[condition] = threshold
		}
	}
}

// validateWeatherCondition 验证天气条件数据
func validateWeatherCondition(c WeatherCondition) error {
	if c.Temperature < -100 || c.Temperature > 100 {
		return &WeatherError{
			Code:    ErrInvalidTemperature,
			Message: fmt.Sprintf("温度数据异常: %.1f°C", c.Temperature),
		}
	}

	if c.Humidity < 0 || c.Humidity > 100 {
		return &WeatherError{
			Code:    ErrInvalidHumidity,
			Message: fmt.Sprintf("湿度数据异常: %.1f%%", c.Humidity),
		}
	}

	if c.WindSpeed < 0 || c.WindSpeed > 100 {
		return &WeatherError{
			Code:    ErrInvalidWindSpeed,
			Message: fmt.Sprintf("风速数据异常: %.1f m/s", c.WindSpeed),
		}
	}

	if c.Precipitation < 0 || c.Precipitation > 1000 {
		return &WeatherError{
			Code:    ErrInvalidPrecipitation,
			Message: fmt.Sprintf("降水量数据异常: %.1f mm", c.Precipitation),
		}
	}

	return nil
}

// Analyze 分析天气状况并返回分析结果
func (wa *WeatherAnalyzer) Analyze() (*WeatherAnalysisResult, error) {
	if len(wa.conditions) == 0 {
		return nil, &WeatherError{
			Code:    ErrEmptyData,
			Message: "没有可分析的天气数据",
		}
	}

	// 计算总降水量和平均降水量
	var totalPrecipitation float64
	var maxPrecipitation float64
	var precipitationHours int
	for _, c := range wa.conditions {
		totalPrecipitation += c.Precipitation
		if c.Precipitation > maxPrecipitation {
			maxPrecipitation = c.Precipitation
		}
		if c.Precipitation > 0 {
			precipitationHours++
		}
	}

	// 计算平均风速和最大风速
	var totalWindSpeed float64
	var maxWindSpeed float64
	for _, c := range wa.conditions {
		totalWindSpeed += c.WindSpeed
		if c.WindSpeed > maxWindSpeed {
			maxWindSpeed = c.WindSpeed
		}
	}
	avgWindSpeed := totalWindSpeed / float64(len(wa.conditions))

	// 根据降水量和风速调整天气状况权重
	adjustedWeights := make(map[string]float64)
	for condition, weight := range wa.conditionWeights {
		adjustedWeights[condition] = weight
	}

	// 根据降水量调整权重
	for condition, threshold := range wa.precipitationThresholds {
		if totalPrecipitation >= threshold {
			// 增加符合降水量条件的天气权重
			adjustedWeights[condition] *= 1.2
		}

	}

	// 根据风速调整权重
	for condition, threshold := range wa.windSpeedThresholds {
		if maxWindSpeed >= threshold {
			// 增加符合风速条件的天气权重
			adjustedWeights[condition] *= 1.15
		}
	}

	// 统计各种天气状况的加权出现次数
	conditionWeightedCount := make(map[string]float64)
	for _, c := range wa.conditions {
		weight := adjustedWeights[c.Condition]
		conditionWeightedCount[c.Condition] += weight
	}

	// 计算平均温度
	var totalTemp float64
	for _, c := range wa.conditions {
		totalTemp += c.Temperature
	}
	avgTemp := totalTemp / float64(len(wa.conditions))

	// 找出加权后出现最多的天气状况
	var maxWeight float64
	var dominantCondition string
	for condition, weight := range conditionWeightedCount {
		if weight > maxWeight {
			maxWeight = weight
			dominantCondition = condition
		}
	}

	// 生成天气描述
	description := fmt.Sprintf("今日天气以%s为主，平均温度%.1f°C", dominantCondition, avgTemp)

	// 添加降水量信息
	if totalPrecipitation > 0 {
		description += fmt.Sprintf("，总降水量%.1f毫米", totalPrecipitation)
		if precipitationHours > 0 {
			description += fmt.Sprintf("，降水持续%d小时", precipitationHours)
		}
		if maxPrecipitation > 0 {
			description += fmt.Sprintf("，最大小时降水量%.1f毫米", maxPrecipitation)
		}
	}

	// 添加风速信息
	if maxWindSpeed > 0 {
		description += fmt.Sprintf("，平均风速%.1f米/秒，最大风速%.1f米/秒", avgWindSpeed, maxWindSpeed)
	}
	description += "。"

	// 添加其他重要天气状况（权重超过总权重的20%）
	var otherConditions []string
	totalWeight := 0.0
	for _, weight := range conditionWeightedCount {
		totalWeight += weight
	}
	threshold := totalWeight * 0.2

	for condition, weight := range conditionWeightedCount {
		if condition != dominantCondition && weight >= threshold {
			otherConditions = append(otherConditions, condition)
		}
	}

	if len(otherConditions) > 0 {
		description += "期间还出现"
		for i, condition := range otherConditions {
			if i > 0 {
				description += "、"
			}
			description += condition
		}
		description += "。"
	}

	// 返回分析结果
	return &WeatherAnalysisResult{
		DominantCondition:  dominantCondition,
		OtherConditions:    otherConditions,
		AverageTemperature: avgTemp,
		TotalPrecipitation: totalPrecipitation,
		MaxPrecipitation:   maxPrecipitation,
		PrecipitationHours: precipitationHours,
		AverageWindSpeed:   avgWindSpeed,
		MaxWindSpeed:       maxWindSpeed,
		ConditionWeights:   conditionWeightedCount,
		Description:        description,
	}, nil
}
