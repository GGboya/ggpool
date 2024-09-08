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
	wg := sync.WaitGroup{}

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
