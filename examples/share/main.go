package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/panjf2000/ants/v2"
)

func Ants() {
	p, _ := ants.NewPool(-1)
	defer p.Release()
	wg := sync.WaitGroup{}
	for i := 2; i <= 100000; i++ {
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

func GoPool() {
	// size 是最大允许并发的数量
	wg := sync.WaitGroup{}

	// 创建任务通道

	// 我们自己写并发,为了控制并发的数目,要限制这个size
	// 只是限制了并发的数量,并没有达到复用的目的,所以速度自然变慢了

	// 提交任务到通道
	for i := 2; i <= 100000; i++ {
		wg.Add(1)
		gopool.Go(func() {
			/// do your job
			defer wg.Done()
			min := 1
			max := 10
			randomDuration := time.Duration(rand.Intn(max-min+1)+min) * time.Nanosecond
			time.Sleep(randomDuration)
		})

	}

	// 关闭任务通道，等待所有任务完成
	wg.Wait()
}

func Gmp() {
	wg := sync.WaitGroup{}
	for i := 2; i <= 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			min := 1
			max := 10
			randomDuration := time.Duration(rand.Intn(max-min+1)+min) * time.Nanosecond
			time.Sleep(randomDuration)

		}(i)

	}
	wg.Wait()

}
