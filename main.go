// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"runtime"
// 	"runtime/trace"
// 	"sync"
// 	"time"
// )

// func main() {
// 	// syncGroupのサンプル
// 	// var wg sync.WaitGroup
// 	// wg.Add(1)
// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	fmt.Println("goroutine invoked")
// 	// }()
// 	// wg.Wait()
// 	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
// 	// fmt.Println("main func finish")
// 	f, err := os.Create("trace.out")
// 	if err != nil {
// 		log.Fatalln("Error:", err)
// 	}
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			log.Fatalln("Error:", err)
// 		}
// 	}()
// 	if err := trace.Start(f); err != nil {
// 		log.Fatalln("Error:", err)
// 	}
// 	defer trace.Stop()
// 	ctx, t := trace.NewTask(context.Background(), "main")
// 	defer t.End()
// 	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

// 	// task(ctx, "task1")
// 	// task(ctx, "task2")
// 	// task(ctx, "task3")
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go cTask(ctx, &wg, "task1")
// 	go cTask(ctx, &wg, "task2")
// 	go cTask(ctx, &wg, "task3")
// 	wg.Wait()
// 	fmt.Println("main func finish")

// }
// // 順次実行のサンプルメソッド
// func task(ctx context.Context, name string) {
// 	defer trace.StartRegion(ctx, name).End()
// 	time.Sleep(time.Second)
// 	fmt.Println(name)
// }
// // 並行実行のサンプルメソッド
// func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
// 	defer trace.StartRegion(ctx, name).End()
// 	defer wg.Done()
// 	time.Sleep(time.Second)
// 	fmt.Println(name)
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// channelのサンプル
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// // leakのサンプル(テストあり)
	// ch1 := make(chan int)
	// go func() {
	// 	fmt.Println(<-ch1)
	// }()
	// ch1 <- 10
	// fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// // bufferのサンプル
	// ch2 := make(chan int, 1)
	// ch2 <- 2
	// fmt.Println(<-ch2)

	// // closeのサンプル
	// ch1 := make(chan int)
	var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println(<-ch1)
	// }()
	// ch1 <- 10
	// close(ch1)
	// v, ok := <-ch1
	// fmt.Printf("%v %v\n", v, ok)
	// wg.Wait()

	// // カプセル化されたチャネルのサンプルを呼び出す
	// ch3 := generateCountStream()
	// for v := range ch3 {
	// 	fmt.Println(v)
	// }

	// 読み取り専用チャネルのサンプル
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v start\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
		time.Sleep(2 * time.Second)
		close(nCh)
		fmt.Println("unlocked by manual close")
		wg.Wait()
		fmt.Println("main goroutine finished")
	}
}

//カプセル化されたチャネルのサンプル
// func generateCountStream() <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 	}()
// 	return ch
// }
