package test

import (
	"testing"
)
const sizeSeq = 1000

func BenchmarkFunc1(b *testing.B) {
	// log.Println(len(seqa), cap(seqa),1)
	for i :=0; i < b.N; i++{
		var seqa []int //len, cap 0
		for i :=0; i < 10; i++{
			seqa = append(seqa, i)
	}
}
	// log.Println(seqa, "res b1 not alloc")
}

//alloc fast minimun 5
func BenchmarkFunc2(b *testing.B) {
// log.Println(len(seq), cap(seq),2)

for i :=0; i < b.N; i++{
	seq  := make([]int, 0, 10) //len 0, cap 1000
	for i :=0; i < 10; i++{
		seq = append(seq, i)
}
}
// log.Println(seq ,"res b2 with alloc")

}
// func TestSequentialFibonacci(t *testing.T) {
//     data := []struct {
//         n    uint
//         want uint
//     }{
//         {0, 0},
//         {1, 1},
//         {2, 1},
//         {3, 2},
//         {4, 3},
//         {5, 5},
//         {6, 8},
//         {10, 55},
//         {42, 267914296},
//     }
//     for _, d := range data {
//         if got := SequentialFibonacci(d.n); got != d.want {
//             t.Errorf("got: %d, want: %d", got, d.want)
//         }
//     }
// }