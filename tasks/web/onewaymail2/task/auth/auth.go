package auth

import (
	"encoding/binary"
	"fmt"
	"gopher/utils"
	"math"
	"net/url"
	"strconv"
)

const SECRET_KEY = "81fa343f4a36a8625d99e523513504fb213920dc713c4d385783528bdf590b727815c5bf"

func KeyStr(id, filename string) (string, error) {
	filename, err := url.QueryUnescape(filename)
	if err != nil {
		return "", err
	}

	c, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	key := make([]byte, 2)
	binary.BigEndian.PutUint16(key, uint16(c%math.MaxUint16))

	key = append(key, []byte(filename)...)
	keyHash := utils.Hash(key)

	keyStr := fmt.Sprintf("%d", binary.LittleEndian.Uint64(keyHash[len(keyHash)-8:]))
	return keyStr, nil
}
