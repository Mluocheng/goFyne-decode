package utils

import (
	"github.com/golang-module/dongle"
)

// 68656c6c6f20776f726c64
func MorseDecode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByMorse("|").ToBytes()
	decodedString := string(decodedBytes)
	return decodedString, nil
}
