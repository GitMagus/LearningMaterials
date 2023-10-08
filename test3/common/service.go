package service

import (
	"math/rand"
	"time"
)

// 100以内的随机数
func GetRandNum() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(100) + 1
}
