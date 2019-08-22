package services

import (
	"fmt"
	"sync"
	"time"
)

func GoBingfa() {
	wg := sync.WaitGroup{}
	//两个channel，一个用来放置工作项，一个用来存放处理结果。
	jobs := make(chan int)
	// 开启三个线程，也就是说线程池中只有3个线程，实际情况下，我们可以根据需要动态增加或减少线程。
	wg.Add(3)
	for w := 1; w <= 3; w++ {
		go func(id int, jobs <-chan int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Println("worker", id, "processing job", j)
				time.Sleep(time.Second)
			}
		}(w, jobs)
	}

	for j := 1; j <= 11; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
