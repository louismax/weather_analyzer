package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func Struct2Base64URL(s interface{}) string {
	str, err := json.Marshal(s)
	if err != nil {
		PrintErrorLog(err.Error())
		return ""
	}
	b64 := base64.RawURLEncoding.EncodeToString([]byte(string(str)))
	return b64
}

func PrintErrorLog(msg string, args ...interface{}) {
	myText := fmt.Sprintf(msg, args...)
	fmt.Printf("\033[31m[--weather_analyzer--]Error:'%s'\033[0m\n", myText)
}

func PrintWarnLog(msg string, args ...interface{}) {
	myText := fmt.Sprintf(msg, args...)
	fmt.Printf("\033[33m[--weather_analyzer--]Warning:'%s'\033[0m\n", myText)
}

func PrintInfoLog(msg string, args ...interface{}) {
	myText := fmt.Sprintf(msg, args...)
	fmt.Printf("\033[34m[--weather_analyzer--]Info:'%s'\033[0m\n", myText)
}
