package digiflazz

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func SignMD5(username string, key string, event string) SignedKey {
	message := fmt.Sprintf("%s%s%s", username, key, event)

	hash := md5.New()
	hash.Write([]byte(message))
	return SignedKey(hex.EncodeToString(hash.Sum(nil)))
}
