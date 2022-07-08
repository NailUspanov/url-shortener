package helpers

const (
	base         uint64 = 63
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
)

func ToBase63(num uint64) string {
	result := ""
	for num > 0 {
		r := num % base
		num /= base
		result = string(characterSet[r]) + result

	}
	return result
}
