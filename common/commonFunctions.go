package common

import (
	"bytes"
	"fmt"
	"strings"
)

/**
*
*author : daozai
*time   : 2017/9/15 22:48
*file   : commonFunctions.go
*
*/

func ConcatBySpace(str ...interface{}) string {
	var buffer bytes.Buffer
	for i := 0; i < len(str); i++ {
		buffer.WriteString(fmt.Sprint(str[i]))
		buffer.WriteString(" ")
	}
	return strings.TrimSpace(buffer.String())
}

func Byte2String(b []byte) string {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return string(b[0:i])
		}
	}
	return string(b)
}