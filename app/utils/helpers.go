package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

func ConvertToJson(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error converting %T to json", v)
		return ""
	}
	return string(b)
}

func MarshalFrom(source any, destination any) error {
	bytes, err := json.Marshal(source)
	json.Unmarshal(bytes, destination)

	return err
}

func ParseInt(valueStr string, defaultValue int) int {
	if valueStr == "" {
		return defaultValue
	}

	if parsedValue, err := strconv.Atoi(valueStr); err == nil && parsedValue > 0 {
		return parsedValue
	}

	return defaultValue
}
