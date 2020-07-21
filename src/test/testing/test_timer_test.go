package testing

import (
	"testing"
	"time"
)

func BenchmarkAdd(b *testing.B) {
	time.Sleep(time.Second)
	b.ResetTimer() //重置记数器，如果执行额外操作应该临时阻止计时器工作
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
		if i == 1 {
			b.StopTimer()
			time.Sleep(time.Second * 3)
			b.StartTimer()
		}
	}
}
