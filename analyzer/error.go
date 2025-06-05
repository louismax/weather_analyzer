package analyzer

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
	ErrInvalidInput         = "INVALID_INPUT"
	ErrEmptyData            = "EMPTY_DATA"
	ErrInvalidTemperature   = "INVALID_TEMPERATURE"
	ErrInvalidHumidity      = "INVALID_HUMIDITY"
	ErrInvalidWindSpeed     = "INVALID_WIND_SPEED"
	ErrInvalidPrecipitation = "INVALID_PRECIPITATION"
)
