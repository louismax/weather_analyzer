package qweather

import "encoding/json"

type ResultQWeatherRefer struct {
	Sources []string `json:"sources"` //原始数据来源，或数据源说明，可能为空
	License []string `json:"license"` //数据许可或版权声明，可能为空
}

type ResultQWeatherError struct {
	Status        int64    `json:"status"`
	Type          string   `json:"type"`
	Title         string   `json:"title"`
	Detail        string   `json:"detail"`
	InvalidParams []string `json:"invalidParams"`
}

type ResultGeoCityLookup struct {
	Code     string                    `json:"code"`
	Location []ResultGeoCityLookupInfo `json:"location"`
	Refer    ResultQWeatherRefer       `json:"refer"`
	Error    ResultQWeatherError       `json:"error"`
}

type ResultGeoCityLookupInfo struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm2      string `json:"adm2"`
	Adm1      string `json:"adm1"`
	Country   string `json:"country"`
	Tz        string `json:"tz"`
	UtcOffset string `json:"utcOffset"`
	IsDst     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

type ResultGeoTopCity struct {
	Code        string                       `json:"code"`
	TopCityList []ResultGeoTopCityListEntity `json:"topCityList"`
	Refer       ResultQWeatherRefer          `json:"refer"`
	Error       ResultQWeatherError          `json:"error"`
}

type ResultGeoTopCityListEntity struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm2      string `json:"adm2"`
	Adm1      string `json:"adm1"`
	Country   string `json:"country"`
	Tz        string `json:"tz"`
	UtcOffset string `json:"utcOffset"`
	IsDst     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

type ResultGeoPoi struct {
	Code  string               `json:"code"`
	Poi   []ResultGeoPoiEntity `json:"poi"`
	Refer ResultQWeatherRefer  `json:"refer"`
	Error ResultQWeatherError  `json:"error"`
}

type ResultGeoPoiEntity struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm2      string `json:"adm2"`
	Adm1      string `json:"adm1"`
	Country   string `json:"country"`
	Tz        string `json:"tz"`
	UtcOffset string `json:"utcOffset"`
	IsDst     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

type ResultQWeatherNow struct {
	Code       string                `json:"code"`
	UpdateTime string                `json:"updateTime"`
	FxLink     string                `json:"fxLink"`
	Now        ResultQWeatherNowInfo `json:"now"`
	Refer      ResultQWeatherRefer   `json:"refer"`
	Error      ResultQWeatherError   `json:"error"`
}

type ResultQWeatherNowInfo struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type ResultQWeatherDaysForecast struct {
	Code       string                `json:"code"`
	UpdateTime string                `json:"updateTime"`
	FxLink     string                `json:"fxLink"`
	Daily      []ResultQWeatherDaily `json:"daily"`
	Refer      ResultQWeatherRefer   `json:"refer"`
	Error      ResultQWeatherError   `json:"error"`
}

type ResultQWeatherDaily struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	Moonset        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	MoonPhaseIcon  string `json:"moonPhaseIcon"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	UvIndex        string `json:"uvIndex"`
}

type ResultQWeatherHistorical struct {
	Code          string                      `json:"code"`
	FxLink        string                      `json:"fxLink"`
	WeatherDaily  ResultWeatherDailyEntity    `json:"weatherDaily"`
	WeatherHourly []ResultWeatherHourlyEntity `json:"weatherHourly"`
	Refer         ResultQWeatherRefer         `json:"refer"`
	Error         ResultQWeatherError         `json:"error"`
}

type ResultWeatherDailyEntity struct {
	Date      string `json:"date"`
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	Moonrise  string `json:"moonrise"`
	Moonset   string `json:"moonset"`
	MoonPhase string `json:"moonPhase"`
	TempMax   string `json:"tempMax"`
	TempMin   string `json:"tempMin"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
}

type ResultWeatherHourlyEntity struct {
	Time      string `json:"time"`
	Temp      string `json:"temp"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Precip    string `json:"precip"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Pressure  string `json:"pressure"`
}

type ResultQWeather struct {
	Body []byte
}

// GeoCityLookupResult GEO城市查询结果解析
func (r *ResultQWeather) GeoCityLookupResult() (*ResultGeoCityLookup, error) {
	result := ResultGeoCityLookup{}
	if err := json.Unmarshal(r.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GeoTopCityResult GEO热门城市查询结果解析
func (r *ResultQWeather) GeoTopCityResult() (*ResultGeoTopCity, error) {
	result := ResultGeoTopCity{}
	if err := json.Unmarshal(r.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GeoPoiResult GEO POI查询结果解析(含poi range)
func (r *ResultQWeather) GeoPoiResult() (*ResultGeoPoi, error) {
	result := ResultGeoPoi{}
	if err := json.Unmarshal(r.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// NowWeatherResult 实时天气查询结果解析
func (r *ResultQWeather) NowWeatherResult() (*ResultQWeatherNow, error) {
	result := ResultQWeatherNow{}
	if err := json.Unmarshal(r.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ForecastWeatherResult 每日预报天气查询结果解析
func (r *ResultQWeather) ForecastWeatherResult() (*ResultQWeatherDaysForecast, error) {
	result := ResultQWeatherDaysForecast{}
	if err := json.Unmarshal(r.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// HistoricalWeatherResult 时光机天气(历史天气)查询结果解析
func (r *ResultQWeather) HistoricalWeatherResult() (*ResultQWeatherHistorical, error) {
	result := ResultQWeatherHistorical{}
	if err := json.Unmarshal(r.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
