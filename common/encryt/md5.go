package encryt

import (
	"crypto/md5"
	"encoding/hex"
)

const defaultSalt = "autowise" // 加盐因子，增加破解的难度

func Md5(encryptStr string) string {
	hash := md5.New()
	hash.Write([]byte(encryptStr + defaultSalt))
	return hex.EncodeToString(hash.Sum(nil))
}
