package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum 将data 转为固定长度为32位的16进制数字
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)) // 32位16进制数
}
