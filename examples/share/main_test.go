package main

import (
	"testing"
)

/*
ants 在百万级别的协程中，内存消耗取得了第一名
不限制协程的数量，让 gmp 自己去调度，内存消耗大概是 ants 的5倍，时间上是 ants 的 2 倍左右
简单的用 channel 来实现对 goroutine 的数量限制，时间上慢， 并且内存消耗也没有显著的优势。
所以 ants 的最大优势我认为是内存的优势，接下来我们去看看他到底是怎么做到自动调度海量goroutine，复用和清理的策略？又是怎么避免内存泄漏的？
*/

func BenchmarkMethod4_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Method4(100000)
	}
}

func BenchmarkMethod4_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Method4(0)
	}
}

func BenchmarkMethod3_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Method3(100000)
	}
}

func BenchmarkMethod3_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Method3(10000)
	}
}

func BenchmarkMethod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Method2()
	}
}
