package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func QueryFromURL[T int | string | []string](param string, r *http.Request) (T, error) {
	queryResult := r.URL.Query().Get(param)
	var conversionResult T

	if queryResult == "" {
		return conversionResult, fmt.Errorf("error while trying to query the parameter: %s from the url", param)
	}

	switch any(conversionResult).(type) {
	case int:
		intValue, err := strconv.Atoi(queryResult)
		if err != nil {
			return conversionResult, fmt.Errorf("error while trying to convert %s to int", queryResult)
		}
		conversionResult = any(intValue).(T)
	case string:
		conversionResult = any(queryResult).(T)
	case []string:
		conversionResult = any(strings.Split(queryResult, ":")).(T)
	default:
		return conversionResult, fmt.Errorf("error while trying to convert %s to a unsuported type", queryResult)
	}

	return conversionResult, nil
}
