package strutl

import (
	"bytes"
	"crypto/rand"
	"github.com/snluu/uuid"
	"strconv"
	"strings"
)

var keywords = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ConnString(strs ...string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func GetUuid() string {
	return strings.ReplaceAll(uuid.Rand().Hex(), "-", "")
}

func GetRandomString(size int) string {
	data := make([]byte, size)
	out := make([]byte, size)
	buffer := len(keywords)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	for id, key := range data {
		x := byte(int(key) % buffer)
		out[id] = keywords[x]
	}
	return string(out)
}

func UrlDecode(s string) string {
	needToChange := false
	numChars := len(s)
	var i int = 0
	var c byte
	result := ""
	var bytes []byte
	for i < numChars {
		c = s[i]
		switch c {
		case '+':
			result += " "
			i++
			needToChange = true
			break
		case '%':
			if bytes == nil {
				bytes = make([]byte, (numChars-i)/3)
			}
			pos := 0
			for (i+2 < numChars) && (c == '%') {
				v, _ := strconv.ParseInt(s[i+1:i+3], 16, 32)
				//                        int v = Integer.parseInt(s.substring(i+1,i+3),16);
				bytes[pos] = byte(v)
				pos++
				i += 3
				if i < numChars {
					c = s[i]
				}
			}

			if (i < numChars) && (c == '%') {

				//                        throw new IllegalArgumentException(
				//                         "URLDecoder: Incomplete trailing escape (%) pattern");
			}

			result += string(bytes[:pos])
			needToChange = true
			break
		default:
			result += string(c)
			i++
			break
		}
	}
	if needToChange {
		return result
	} else {
		return s
	}
}