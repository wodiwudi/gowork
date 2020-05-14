package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//任务的取消
//使用context上下文来实现，上一个channel取消任务 只能是叶子结点，
//如果为父任务，则不会连带取消子任务改造为使用context
func isContextcancelChan(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestContextCancelChan(t *testing.T) {
	//context.Background()为创造context根结点
	//context.withCancel为创造一个子结点，第二个参数为一个cancel当前子结点的取消方法。
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		//创造5个协程
		go func(num int, ctx context.Context) {
			for {
				if isContextcancelChan(ctx) {
					break
				}
				time.Sleep(time.Second * 1)
			}
			fmt.Printf("%d is done !", num)
			time.Sleep(time.Second * 1)
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
