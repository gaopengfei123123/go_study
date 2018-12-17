package main

import (
	"testing"
)

func TestDoSomething(t *testing.T) {
	record := doSomething()

	t.Log(record)
}

func BenchmarkDoSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doSomething()
	}
}

func BenchmarkDoSomethingParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doSomething()
		}
	})
}

func BenchmarkDoSomethingWithoutPoll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doSomething2()
	}
}

func BenchmarkDoSomethingWithoutPollParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doSomething2()
		}
	})
}
