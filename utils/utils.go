package utils

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

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

// GenerateRandomString generates a random string of the given length.
func GenerateRandomString(length int) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}
