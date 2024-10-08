package utils

import (
	"math/rand"
	"time"
)

// ASCII values 33 ~ 126
const _dcl = 126 - 33 + 1

var defaultCharacters [_dcl]byte

func init() {
	for i := 0; i < _dcl; i++ {
		defaultCharacters[i] = byte(i + 33)
	}
}

func RandString(length int, characters ...byte) string {
	if len(characters) == 0 {
		characters = defaultCharacters[:]
	}

	n := len(characters)
	var rs = make([]byte, length)

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		rs[i] = characters[randGen.Intn(n-1)]
	}

	return string(rs)
}
