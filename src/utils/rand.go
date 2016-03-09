package utils

import (
	"math/rand"
	"time"
)

var rander *rand.Rand

func InitRander() {
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRand() *rand.Rand {
	return rander
}
