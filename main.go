package main

import (
	"fmt"
	"sync"
)

var i int
var mu sync.Mutex
// 待ち合わせ
var wg sync.WaitGroup

func MyFunc1() {
	// 内部カウンタを-1
	defer wg.Done()

	fmt.Println("MyFunc1 start")	
	
	mu.Lock()
	i += 1
	mu.Unlock()
	
	fmt.Println("MyFunc1 finish")
}

func MyFunc2() {
	// 内部カウンタを-1
	defer wg.Done()

	fmt.Println("MyFunc2 start")

	mu.Lock()
	i -= 1
	mu.Unlock()

	fmt.Println("MyFunc2 finish")
}

func main() {
	i = 0
	
	// 内部カウンタを+2にする
	wg.Add(2)
	go MyFunc1()
	go MyFunc2()

	// 内部カウンタが0になるまで待機
	wg.Wait()
	fmt.Printf("final result: i = %d\n", i)
}