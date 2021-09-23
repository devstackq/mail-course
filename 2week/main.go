package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

//context

func main() {
	var count int32
	//race condition, all goroutine 1 time - try change value
	// mu := &sync.Mutex{}
	//trye use - mutext or atomic ?
	wg := sync.WaitGroup{}
	wg.Add(1000)

	//atomic

	for i := 0; i < 1000; i++ {
		go func() {
			// mu.Lock()
			// count++
			// mu.Unlock()
			atomic.AddInt32(&count, 2)
			wg.Done()

		}()
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
	//time - for call
	log.Println(count, "count")
	//1 case:
	// ctx, cancel := context.WithCancel(context.Background())

	// ch := make(chan int)

	// for i := 1; i < 10; i++ {
	// 	go run(ctx, ch, i)
	// }

	// //read from chan
	// for i := 1; i < 10; i++ {
	// 	n := <-ch
	// 	//write from chan, compare received data, if cond = true, cancel, break loop
	// 	if n > i && n%10 == 0 {
	// 		log.Println("cond true, call cancel() ctx", n)
	// 		cancel()
	// 		break
	// 	}
	// 	log.Println(n, "read chan")
	// }

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)

	// log.Println(runtime.NumGoroutine(), "num go")
	//chan receive data ? call cancel()

	//set data from func
	// chanData := getComments()

	// time.Sleep(1 * time.Second)

	// log.Println("get articles func")

	// //read from chan
	// return
	// d := <-chanData
	// log.Println("mainGo read data from gorotuine called func data", d)
	// getPage()

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		//scheduler - run 3 gorutine worker(); each workegr,/
		//mainGo - loop: add jobsChan <- val, concurrent worker() complete func,
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		log.Println(<-results, "mainGo read chan result")
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	//shceduler manage gorotine, queuee -> like fifo
	for j := range jobs { //range queue int value from chan- mainGo values chan[int] buff chan
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func getPage() {
	resultCh := getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")
	//case2
	// return

	commentsData := <-resultCh
	fmt.Println("main goroutine:", commentsData)
}
func getComments() chan string {

	ch := make(chan string, 1) // if chan no buff - утечка горутин и памяти
	//garbage collector- сли никто не записывает и не читает с канала

	go func(ch chan string) {
		log.Print("get from db data")
		time.Sleep(2 * time.Second)
		ch <- "33 comments"
		time.Sleep(1 * time.Second) // case 2 , утечка горутин, канал переполнен, Ю блокировка записи , тюк никто не читает
		ch <- "4 comments"
	}(ch)
	return ch
}

func run(ctx context.Context, ch chan int, iter int) {
	ch <- iter * 2
	log.Println(iter, "go iter num")
	select {
	case <-ctx.Done():
		log.Println("cancel(), ctx.Done, drop goroutines")
		return
	}
}
