package Singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
	data int
}

var sin *singleton

var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() {
		sin = &singleton{data: 12}
	})
	fmt.Println("实例对象", sin, &sin)
	return sin
}
