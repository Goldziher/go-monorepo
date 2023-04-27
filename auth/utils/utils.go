package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

func CreateState() (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, 15)
	for i := 0; i < 15; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return base64.URLEncoding.EncodeToString(ret), nil
}
