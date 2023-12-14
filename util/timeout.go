package util

import (
	"fmt"
	"time"
)

func WorkWithTimeout(timeout time.Duration) {
	workCh := make(chan struct{}, 1)
	go func() {
		//LongTimeWork() //把要控制超时的函数放到子协程里去执行
		workCh <- struct{}{}
	}()
	select { //只执行最先到来的case
	case <-workCh: //work先结束
		fmt.Println("work finish")
	case <-time.After(timeout): //超时先来
		fmt.Println("work timeout")
	}
}
