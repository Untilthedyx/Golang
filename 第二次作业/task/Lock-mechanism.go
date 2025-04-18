package task

import (
	"sync"
	"sync/atomic"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
var Mu sync.Mutex
var count int32 = 0
var Mutex5, Mutex6 sync.WaitGroup

func Mutex() int32 {

	for i := 0; i < 10; i++ {
		Mutex5.Add(1)
		go func() {
			Mu.Lock()
			for j := 0; j < 1000; j++ {
				count++
			}
			Mu.Unlock()
			Mutex5.Done()
		}()

	}
	Mutex5.Wait()
	return count
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func Atomic() int32 {
	for i := 0; i < 10; i++ {
		Mutex6.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
			Mutex6.Done()
		}()
	}
	Mutex6.Wait()
	return count
}
