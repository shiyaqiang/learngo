package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	// lock 是一个全局互斥锁 sync 是包 同步的意思 Mutex:是互斥
	lock sync.Mutex
)
// test函数计算 n的阶乘, 将结果放到map中
func testCount(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *=i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}
func main () {
	// 开启多个协程完成20个任务
	a :=time.Now()

	for i := 1; i <= 19; i++{
		testCount(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for k, v := range myMap{
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()
	e :=time.Since(a)
	fmt.Println(e)

}