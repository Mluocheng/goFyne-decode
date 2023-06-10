// 字符集 解码
package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/axgle/mahonia"
)

// C4E3BAC3D1BDA3ACBAC3BFAAD0C4A3A1
func Ansi2String(ascii string) string {
	var baKeyword = make([]byte, len(ascii)/2)
	for i := 0; i < len(baKeyword); i++ {
		baKeyword[i] = byte(0xff & StringToInt64(ascii[i*2:i*2+2]))
	}
	return mahonia.NewDecoder("GBK").ConvertString(string(baKeyword[:]))
}

func StringToInt64(str string) int64 {
	// fmt.Println("str:", str)
	str = strings.ReplaceAll(str, "\r", "") // 去除换行符 \r
	str = strings.ReplaceAll(str, "\n", "") // 去除换行符 \n
	str = strings.ReplaceAll(str, "\\", "") // 去除换行符

	if val, err := strconv.ParseInt(str, 16, 64); err != nil {
		// panic(err)
	} else {
		return val
	}
	return 0
}

// 4F60597D5440FF0C597D5F005FC3FF01
func Ucs2Tostring(str string) string {
	if len(str)%4 != 0 || ContainsSpecialChar(str) {
		return ""
	}

	var sUnicodev = make([]string, 0)
	for i := 0; i < len(str); i = i + 4 {
		stri := str[i : i+4]
		if len(stri) > 0 {
			sUnicodev = append(sUnicodev, stri)
		}
	}
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			// panic(err)
			return ""
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}

// E4BDA0E5A5BDE59180EFBC8CE5A5BDE5BC80E5BF83EFBC81
func Utf8Encode2String(str string) string {
	hexData, err := hex.DecodeString(str)
	if err != nil {
		// panic(err)
		return ""
	}
	// 将 byte 转换 为字符串 输出结果
	return string(hexData)
}

// \u0048\u0065\u006c\u006c\u006f\u002c\u0020\u0057\u006f\u0072\u006c\u0064\u0021
func Utf16Encode2String(str string) string {
	// fmt.Println("str11111111:", str, "str11111111:")
	// 解析转义序列为 Unicode 字符
	runes := make([]rune, 0)
	parts := strings.Split(str, "\\u")
	for _, part := range parts {
		if part == "" {
			continue
		}
		code, _ := strconv.Unquote(`"\u` + part + `"`)
		runes = append(runes, []rune(code)...)
	}

	// UTF-16 解码
	uint16s := utf16.Encode(runes)
	decoded := string(utf16.Decode(uint16s))

	return decoded
}

// "\x00\x00\x00H\x00\x00\x00e\x00\x00\x00l\x00\x00\x00l\x00\x00\x00o\x00\x00\x00,\x00\x00\x00 \x00\x00\x00W\x00\x00\x00o\x00\x00\x00r\x00\x00\x00l\x00\x00\x00d\x00\x00\x00!"
func UTF32ToString(encoded string) (string, error) {
	encoded = strings.Trim(encoded, "\n")
	fmt.Println("encoded:", encoded)
	// 检查输入字符串长度是否为 UTF-32 编码的倍数
	if len(encoded)%4 != 0 {
		return "", fmt.Errorf("输入字符串长度不是 UTF-32 编码的倍数")
	}

	// 将字符串转换为字节数组
	var bytes []byte
	for i := 0; i < len(encoded); i += 4 {
		val := encoded[i : i+4]
		b, err := decodeHex(val)
		if err != nil {
			return "", fmt.Errorf("解析字节失败：%v", err)
		}
		bytes = append(bytes, b)
	}

	// 创建一个切片来存储解码后的 uint32 值
	uint32s := make([]uint32, len(bytes)/4)

	// 逐个读取 4 个字节并解码为 uint32 值
	for i := 0; i < len(uint32s); i++ {
		val := binary.BigEndian.Uint32(bytes[i*4 : (i+1)*4])
		uint32s[i] = val
	}

	// 将 uint32 数组转换为字符串
	runes := make([]rune, len(uint32s))
	for i, u := range uint32s {
		runes[i] = rune(u)
	}

	return string(runes), nil
}
func decodeHex(val string) (byte, error) {
	var b byte
	for i := 0; i < len(val); i += 2 {
		var digit byte
		switch {
		case val[i] >= '0' && val[i] <= '9':
			digit = val[i] - '0'
		case val[i] >= 'a' && val[i] <= 'f':
			digit = val[i] - 'a' + 10
		case val[i] >= 'A' && val[i] <= 'F':
			digit = val[i] - 'A' + 10
		default:
			return 0, fmt.Errorf("非法的十六进制字符：%c", val[i])
		}
		b = (b << 4) | digit
	}
	return b, nil
}
