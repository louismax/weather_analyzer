package utils

import "fmt"

// WeatherError 自定义天气分析错误类型
type WeatherError struct {
	Code    string
	Message string
	Err     error
}

func (e *WeatherError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// 预定义错误码
const (
	ErrInvalidInput         = "INVALID_INPUT"         // 输入无效
	ErrEmptyData            = "EMPTY_DATA"            // 数据为空
	ErrInvalidTemperature   = "INVALID_TEMPERATURE"   // 温度无效
	ErrInvalidHumidity      = "INVALID_HUMIDITY"      // 湿度无效
	ErrInvalidWindSpeed     = "INVALID_WIND_SPEED"    // 风速无效
	ErrInvalidPrecipitation = "INVALID_PRECIPITATION" // 降雨量无效
	ErrReadFile             = "READ_FILE_ERROR"       // 读取文件错误
	ErrPrivateKeyInvalid    = "PRIVATE_KEY_INVALID"   // 私钥无效
	ErrRequestFailed        = "REQUEST_FAILED"        // 请求失败
)
