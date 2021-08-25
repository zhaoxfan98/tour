package util

import (
	"crypto/md5"
	"encoding/hex"
)

//用于针对上传后的文件名格式化 避免暴露原始名称
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
