package utils

// TODO: move to the corresponding resource
import (
	"crypto/md5"
	"encoding/hex"
)

func HashStringMD5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
