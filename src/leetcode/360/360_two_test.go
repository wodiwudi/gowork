package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test360Two(t *testing.T) {
	getit(3, 3)
}
func getit(x, y int) {
	fmt.Println(y)
	fmt.Println(rand.Intn(y))
}
