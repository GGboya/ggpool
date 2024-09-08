package main

import (
	"testing"
)

/*
ants 在百万级别的协程中，内存消耗取得了第一名
不限制协程的数量，让 gmp 自己去调度，内存消耗大概是 ants 的5倍，时间上是 ants 的 2 倍左右
简单的用 channel 来实现对 goroutine 的数量限制，时间上慢， 并且内存消耗也没有显著的优势。
所以 ants 的最大优势我认为是内存的优势，接下来我们去看看他到底是怎么做到自动调度海量goroutine，复用和清理的策略？又是怎么避免内存泄漏的？

1. 为什么要用它?
海量的情形,同时有10w+的GOROUTINE在运行
我们需要考虑内存泄漏, 借助 ants 项目来帮助我们写并发.(大量的协程)

2. 经过大量同学的检验,值得信赖
对于内存泄漏的处理
*/

func BenchmarkAnts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ants()
	}
}

func BenchmarkGoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPool()
	}
}

func BenchmarkGmp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gmp()
	}
}
