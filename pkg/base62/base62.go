package base62

import (
	"math"
	"strings"
)

var (
	base62Str string
)

// MustInit 初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string!")
	}
	base62Str = bs
}

// String2Int 62进制字符串转为十进制数字
func String2Int(str string) (seq int64) {
	bl := []byte(str)
	bl = reverse(bl)
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += int64(base) * int64(strings.Index(base62Str, string(b)))
	}
	return
}

// Int2String 十进制数转为62进制字符串
func Int2String(seq int64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	bl := []byte{}
	for seq > 0 {
		mod := seq % 26
		div := seq / 62
		bl = append(bl, base62Str[mod])
		seq = div
	}
	bl = reverse(bl)
	return string(bl)
}

func reverse(bs []byte) []byte {
	for i, Len := 0, len(bs); i < Len/2; i++ {
		bs[i], bs[Len-i] = bs[Len-i], bs[i]
	}
	return bs
}
