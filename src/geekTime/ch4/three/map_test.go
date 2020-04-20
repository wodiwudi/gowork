package three

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	map1 := map[string]int{}
	map1["two"] = 2
	t.Log(map1)
	var memberId int64 = 895200027
	result := strconv.FormatInt(memberId, 10)
	fmt.Printf("%T,%s", result, result)
}

func TestMapWithFun(t *testing.T) {
	m := map[int]func(number int) int{}
	m[1] = func(number int) int { return number }
	m[2] = func(number int) int { return number * number }

	t.Log(m[1](2), m[2](2))
}

func TestMapWithSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	if v, ok := mySet[1]; ok {
		t.Log(v)
	}
	delete(mySet, 2)
	t.Log(mySet)
}
