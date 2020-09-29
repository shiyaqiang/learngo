package main

import "fmt"

// 往管道里写入50条数据
func writeDate(intChan chan int)  {
	for i := 1; i <= 200; i++{
		// 写入数据
		intChan<- i
		fmt.Println("管道中写入数据:", i)
	}
	// 写完后,关闭此管道
	close(intChan)
}
// 从管道中读取数据
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		// 说明intChan管道已经取完了
		if !ok{
			break
		}
		fmt.Printf("intChan 管道取出数据:%v\n", v)
	}
	// readData 取完后表示任务已经完成
	exitChan<- true
	close(exitChan)
}
func main()  {
	// 创建两个管道
	intChan := make(chan int, 2)
	// 退出管道, 主线程监控, 协程取完intChan后, 会写进此管道一条数据
	exitChan := make(chan bool, 1)

	// 开启写的协程、读的协程
	go writeDate(intChan)
	go readData(intChan, exitChan)

	// 写一个for循环, 监听exitChan管道, 若exitChan管道的数据取完, 主线程可以结束
	// for {
	// 	_, ok := <- exitChan
	// 	if !ok{
	// 		break
	// 	}
	// }
	<- exitChan
}