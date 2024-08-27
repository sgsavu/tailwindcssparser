package tailwindcssparser

import (
	"time"
	"unsafe"

	"golang.org/x/exp/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 64 / letterIdxBits
)

var src = rand.NewSource(uint64(time.Now().UnixNano()))

func getRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Uint64(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Uint64(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
