package utils

import "math"

func GetCode(number int64, length int64) string {
	return GetCodeFromTemplate(number, 3, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func GetCodeFromTemplate(number int64, length int64, template string) string {
	var result = ""
	var i = int64(0)
	var ln = int64(len(template))
	var capacity = int64(math.Pow(float64(ln), float64(length)))
	number = number % capacity
	for i < length {
		var cur = number % ln
		if i > 0 {
			cur = (cur + int64(result[i-1])) % ln
		}
		result = result + string(template[cur])
		number = number / ln
		i++
	}
	return result
}
