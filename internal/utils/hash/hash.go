package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func ConvertToHash(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}
