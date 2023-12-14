package chang

import (
	"fmt"
	"sync"
	"time"
)

/*
主程序,之前试过让最后一个协程关闭msg通道,但是最后一个协程并不知道一起并发的另外四个进程是否结束,有可能丢数据
现在存在的问题
1.使用延时方法关闭channel具有不确定性,跑真实数据的时候不确定什么时候正常结束退出
2.使用了据说不推荐的goto LABEL 方法
*/

//func msaster() {
//	limit := make(chan struct{}, 5)
//	msg := make(chan string, 1)
//	go func() {
//		for i := 0; i < 50; i++ {
//			limit <- struct{}{}
//			go mission(i, msg, limit)
//		}
//	}()
//	//用来保存从所有通道获取的结果
//	var loop = true
//	var finally []string
//	for loop {
//		select {
//		case data := <-msg:
//			finally = append(finally, data)
//		case <-time.After(5 * time.Second):
//			fmt.Println("Timeout: No data received,after 3 second")
//			loop = false
//
//			fmt.Printf("finally is %v\ntype is %T\n", finally, finally)
//		}

func mission(index int, msg chan string, limit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	// 这是一个长时间任务
	for i := 0; i < 10; i++ {
		fmt.Printf("这是第%d个协程的第%d次响应\n", index+1, i)
		time.Sleep(300 * time.Millisecond)
	}
	//实际上msg需要传输经过函数处理的内容
	msg <- fmt.Sprintf("这是第%d个协程\n", index)
	<-limit
	fmt.Printf("从通道中释放一个空结构体\n")
}

func master() {
	limit := make(chan struct{}, 5)
	msg := make(chan string, 50)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		limit <- struct{}{}
		wg.Add(1)
		go mission(i, msg, limit, &wg)
	}

	go func() {
		wg.Wait()
		close(msg)
	}()

	// 用来保存从所有通道获取的结果
	var finally []string
	for data := range msg {
		fmt.Printf("从msg通道获取的msg:%v\n", data)
		finally = append(finally, data)
	}

	fmt.Printf("finally is %v\ntype is %T\n", finally, finally)
}
