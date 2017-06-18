package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jmcvetta/randutil"
)

func RandomString() string {
	return RandomStringWithLength(12)
}

func RandomStringWithLength(len int) string {
	rand, _ := randutil.String(12, randutil.Alphabet)
	ts := time.Now().Unix()
	s := TruncateString(fmt.Sprintf("%v%v", rand, ts), len)
	return s
}

func RandomIntAsString(min, max int) string {
	return strconv.Itoa(RandomInt(min, max))
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func TruncateString(s string, i int) string {
	if len(s) < i {
		return s
	}
	return s[:i]
}
