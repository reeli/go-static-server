package main

import (
	"fmt"
	"time"
)

// 1. 类型
// 2. go build / go run

func main() {

	var i int

	//rw := &sync.RWMutex{}

	for j := 0; j < 5; j++ {
		//1. 变量赋值, 垃圾回收机制？
		//2. 函数传参

		n := j

		go func(a int) {
			//time.Sleep(50*time.Millisecond)
			//rw.Lock()
			//defer rw.Unlock()

			i = n

			fmt.Println(i, n)
		}(j)
	}

	time.Sleep(50 * time.Millisecond)

	fmt.Println(i)
	//server.Serve()
}

// read 加锁，保证在同一个函数执行的时候，变量的值不会被其他线程修改了
