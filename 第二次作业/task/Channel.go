package task

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
var Mutex3, Mutex4 sync.WaitGroup

func Channel1() {
	Mutex3.Add(2)
	ch := make(chan int)
	go func(chan<- int) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
		Mutex3.Done()
	}(ch)
	go func(<-chan int) {
		for i := range ch {
			fmt.Printf("%d", i)
		}
		Mutex3.Done()
	}(ch)
	Mutex3.Wait()

}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

var ch chan int = make(chan int, 100)

func Producer(hundred [100]int) {
	for i := range hundred {
		ch <- i
	}
	close(ch)
	Mutex4.Done()
}

func Consumer() {
	for i := range ch {
		fmt.Printf("%d", i)
	}
	Mutex4.Done()
}
func Channel2() {
	Mutex4.Add(2)
	go Producer([100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	go Consumer()
	Mutex4.Wait()
}
