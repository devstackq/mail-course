package main

import "log"


func main(){

	// ch := make(chan int)
// ctx := context.Background()//context main

// var wg sync.WaitGroup
// wg.Add(3)
// ch <- 22
// ch <- -2
// ch <- 33

// 	go func(ctx context.Context,c chan int){
// 		n := <- ch
// 		log.Println("val1", n)
// 		if  n < 0 {
// 		log.Println("drop gorotuine1 <", n)
// 			cancel()
// 		}
// 		wg.Done()
// 	}(ctx, ch)

// 	go func(ctx context.Context,c chan int){
// 		// log.Println(<-c)
// 		n := <- ch
// 		log.Println("val2", n)
// 		if n > 0 {
// 		log.Println("drop gorotuine2 >", n)
// 			cancel()
// 		}
// 		wg.Done()
// 	}(ctx, ch)


// 	go func(ctx context.Context,c chan int){
// 		// log.Println(<-c)
// 		n := <- ch
// 		log.Println("val3", n)
// 		if n > 0 {
// 		log.Println("drop gorotuine3 >", n)
// 			cancel()
// 		}
// 		wg.Done()
// 	}(ctx, ch)

// ch := make(chan int)
// // ctx, cancel := context.WithCancel(context.Background())

// //run gorutine, case 1: chan read
// 	go func(out chan int ) {
// 		for i :=0; i <=4; i++ {
// 			log.Println("bef", i)
// 	out <- i//write chan
// 	log.Println("aft", i)
// 	// break
// }
// 		// close(out)// if close, loop - end
// 		out <- 10
// 		//then close chan
// 	}(ch)

// //sequence execute, write, read in cahnnel data
// // for i :=0; i <1; i++ {
// 	for i:= range ch {
// 	log.Println(i, "mainG read chan")//read chan
// }


//multiplexor

// ch1 := make(chan string,1) //buf
// ch2 := make(chan string,2)

// //read/write ion chan


// ch1 <-"a"
// ch2<- "b"
// ch2<- "z"
// // ch2 <-"c"

// //loop check -> each channel loop
// //sheduler -> sam vybirate kakoi pervyi kanal orichest
// LOOPA:
// for {
// select {
// case val := <- ch1:
// 	log.Println(val, "read data from chan 1")
// case  v2 := <- ch2:
// 	log.Println(v2, "put data from chan 2")
// default:
// 	log.Println("def case, break")
// 	break  LOOPA// break beacuse, chan 
// }
// }
//if !def case - read/write no in channels -> deadlock 


chData := make(chan int)
cancelChan := make(chan struct{})

go func(data chan int, cancel chan struct{}) {
val :=0
	for {
		select {
		case  <- cancel:
			log.Print("cancel2")
			return
		case chData  <- val: //0,1,2,3,.break
		log.Print(val,"ss")
			val++
		}
	}

}(chData, cancelChan)


//read from goroutine, incremented chData
for val := range chData {
	if val >3 {
		cancelChan <- struct{}{} // epmty struct
		break
	}else {
		log.Println("data", val)
	}
}

}

//
