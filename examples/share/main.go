package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	fmt.Println("执行开始")
	fmt.Println("完成")
}

type task func()

func Method4(size int) {
	p, _ := ants.NewPool(size)
	defer p.Release()
	wg := sync.WaitGroup{}
	for i := 2; i <= 1000000; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			defer wg.Done()
			min := 1
			max := 10
			randomDuration := time.Duration(rand.Intn(max-min+1)+min) * time.Nanosecond
			time.Sleep(randomDuration)
		})
	}
	wg.Wait()
}

func Method3(size int) {
	// size 是最大允许并发的数量
	wg := sync.WaitGroup{}

	// 创建任务通道
	limit := make(chan struct{}, size)

	// 提交任务到通道
	for i := 2; i <= 1000000; i++ {
		wg.Add(1)
		limit <- struct{}{} // 放入一个空结构体，阻塞时表示协程数达到上限
		go func() {
			defer wg.Done()
			defer func() { <-limit }() // 完成后释放一个空结构体
			min := 1
			max := 10
			randomDuration := time.Duration(rand.Intn(max-min+1)+min) * time.Nanosecond
			time.Sleep(randomDuration)
		}()
	}

	// 关闭任务通道，等待所有任务完成
	wg.Wait()
}

func Method2() {
	wg := sync.WaitGroup{}
	for i := 2; i <= 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 生成 10 毫秒到 100 毫秒之间的随机时间
			min := 1
			max := 10
			randomDuration := time.Duration(rand.Intn(max-min+1)+min) * time.Nanosecond
			time.Sleep(randomDuration)

		}(i)

	}
	wg.Wait()

}

// func Method1() {
// 	for i := 2; i <= 10000; i++ {
// 		// 生成 10 毫秒到 100 毫秒之间的随机时间
// 		min := 1
// 		max := 100
// 		randomDuration := time.Duration(rand.Intn(max-min+1)+min) * time.Nanosecond
// 		time.Sleep(randomDuration)
// 	}
// }
