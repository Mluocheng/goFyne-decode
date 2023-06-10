package utils

import (
	"encoding/base64"

	"github.com/golang-module/dongle"
)

// 68656c6c6f20776f726c64
func Base16Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase16().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// NBSWY3DPEB3W64TMMQ======
func Base32Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase32().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// +8D VD82EK4F.KEA2
func Base45Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase45().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// StV1DL6CwTryKyV
func Base58Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase58().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// AAwf93rvy4aWQVw
func Base62Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase62().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// aGVsbG8gd29ybGQ=
func Base64Decode(encodedString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", err
	}

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// d3d3LmdvdWd1b3lpbi5jbg==
func Base64UrlDecode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase64URL().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// BOu!rD]j7BEbo7
func Base85Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase85().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}

// TPwJh>Io2Tv!lE
func Base91Decode(encodedString string) (string, error) {
	var decodedBytes = dongle.Decode.FromBytes([]byte(encodedString)).ByBase91().ToBytes()

	decodedString := string(decodedBytes)
	return decodedString, nil
}
