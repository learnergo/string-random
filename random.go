package string_random

import (
	"math"
	"math/rand"
)

const LOWERLETTERS = "abcdefghijklmnopqrstuvwxyz"
const UPPERLETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DIGIT = "0123456789"
const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALL = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Rand(n int) string {
	return innerRand(ALL, n)
}

func UpperLetterRand(n int) string {
	return innerRand(UPPERLETTERS, n)
}

func LowerLetterRand(n int) string {
	return innerRand(LOWERLETTERS, n)
}

func DigitRand(n int) string {
	return innerRand(DIGIT, n)
}

func innerRand(letterBytes string, n int) string {
	length := len(letterBytes)
	letterIdxBits := uint(math.Ceil(math.Log2(float64(length))))
	var (
		letterIdxMask int64 = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax        = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

type random struct {
	scope string
}

func (r *random) Rand(n int) string {
	return innerRand(r.scope, n)
}

func New(scope string) *random {
	return &random{
		scope: scope,
	}
}
