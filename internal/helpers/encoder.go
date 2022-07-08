package helpers

import (
	"encoding/hex"
	"strconv"
	"strings"
)

const encodingSize = 2
const decodingSize = 3

const (
	base         uint64 = 63
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
)

func Encode(str string) string {
	var result strings.Builder

	bytes := []byte(str)
	bytesLength := len(bytes)

	// преобразуем строки в десятичное представление, переведя байты сначала в 16ю, после в 10ю СС
	// так как в uint64 может возникнуть переполнение, кодировать будем поочередно (по 2 байта)
	// соответственно, если максимальное значение 2Б для 16 СС = ffff => 65535 в 10 СС => GWF в нашей кодировке
	// следовательно, каждая итерация кодирования может вернуть максимум 3 символа кодировки
	// поэтому нужно учесть случаи, когда после кодирование возвратится <3 символов кодировки
	// (чтобы можно было легко раскодировать полученную строку с помощью обратного алгоритма)
	// для этого в начало будем добавлять незначащие нули
	for i := 0; i < bytesLength; i += encodingSize {
		substringBytes := bytes[i:min(i+encodingSize, bytesLength)]
		h := hex.EncodeToString(substringBytes)
		val, _ := strconv.ParseUint(h, 16, 64)
		w := padLeft(toBase63(val), "0", decodingSize)
		result.WriteString(w)
	}
	return result.String()[0:10]
}

func toBase63(num uint64) string {
	result := ""
	for num > 0 {
		r := num % base
		num /= base
		result = string(characterSet[r]) + result

	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func padLeft(str, pad string, length int) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}
