package gmgn

import (
	"strconv"
)

func GetIntValue(value interface{}) int {
	valueStr, _ := value.(string)
	intValue, _ := strconv.Atoi(valueStr)
	return intValue
}

func GetFloatValue(value interface{}) float64 {
	valueStr, _ := value.(string)
	floatValue, _ := strconv.ParseFloat(valueStr, 64)
	return floatValue
}
