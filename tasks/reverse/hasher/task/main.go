package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

var FLAG_CONTENT = os.Getenv("FLAG_CONTENT")

func Sha512(bytes []byte) string {
	h := sha512.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func MeasureExecution(runtime func()) float64 {
	startTime := time.Now().UnixNano()
	runtime()
	endTime := time.Now().UnixNano()
	seconds := (float64(endTime) - float64(startTime)) / float64(time.Second)
	return seconds
}

func main() {
	var cmd string
	fmt.Scan(&cmd)
	block := Sha512([]byte(cmd))
	flag_chars := []rune(FLAG_CONTENT)
	executed := MeasureExecution(func() {
		for i, char := range []rune(cmd) {
			if string(flag_chars[i]) == string(char) {
				for j := 0; j < 100; j++ {
					block = Sha512([]byte(block + string(char)))
				}
			}
		}
	})
	fmt.Println(fmt.Sprintf("sha512: %s %f", block, executed))
}
