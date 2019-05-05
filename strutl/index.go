package strutl

import (
	"bytes"
	"crypto/rand"
)

var keywords = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ConnString(strs ...string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
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