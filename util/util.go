package util

import (
	"math/rand"
	"time"
)

//Code to build 'random' strings inspired by:
//https://stackoverflow.com/a/31832326/448591

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randSeeded  = false
)

func RandString(n int) string {
	if n <= 0 {
		return ""
	}
	if !randSeeded {
		rand.Seed(time.Now().UnixNano())
		randSeeded = true
	}
	//We got stuff to do:
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
