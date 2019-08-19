package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateOTP() (string, error){
	nBig, e := rand.Int(rand.Reader, big.NewInt(89999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+10000, 10), nil
}