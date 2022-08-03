package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

func Md5Encrypt(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func RandomString(n int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	s := make([]byte, n)
	for i := range s {
		s[i] = chars[r.Intn(len(chars))]
	}
	return string(s)
}
