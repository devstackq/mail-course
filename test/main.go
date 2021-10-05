package main

func main() {

}

// var cache map[int]int


func fib(n int)int {
	if n<=1 {
		return n
	}
	return  fib(n-1)+fib(n-2)
}

func FindPrime(n int) bool{

	for i := 2; i < n ; i++{
	 if n % i == 0 {
	  return false
	}
	}
	return true
	}