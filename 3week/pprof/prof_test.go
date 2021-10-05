package prof

import (
	"io/ioutil"
	"log"
	"testing"
)

//test
func TestSolve( t *testing.T) {
	c, _ := ioutil.ReadFile("./note.txt")
	res := Solve(c)
	log.Println(res,"test res")
	excepted := 3
	if excepted != res {
		t.Fatalf("excepted %d, got %d", excepted, res)
	}
}

//bench
func BenchSolve( b *testing.B) {
	c, _ := ioutil.ReadFile("./note.txt")
	for i :=0; i < b.N; i++{
		Solve(c)
	}
}


func TestPrime(t *testing.T) {

	data := []struct {
		        n    int
		        want bool
		    }{
		        {10, false},
		        {3, true},
		        {20, false},
		        {13, true},
		        {4, false},
		        {52, false},
		    }

		    for _, d := range data {
		        if got := FindPrime(d.n); got != d.want {
		            t.Errorf("got: %t, want: %t, val: %d", got, d.want, d.n)
		        }else {
					t.Log("Oks", d.n)
				}
		    }
}