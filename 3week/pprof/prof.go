package prof

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
)

// Each line gives the password policy and then the password.
// The password policy indicates the lowest and highest number of times
// a given letter must appear for the password to be valid.

// 1-3 a: aabacdlog.Println(string(contents))

func Solve(contents []byte) int {
	valid := 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		lineStr := string(line)
		layout := regexp.MustCompile(`([0-9]+)-([0-9]+)\s(\w){1}:\s([a-z]+)$`)
		matches := layout.FindStringSubmatch(lineStr)
		if len(matches) < 5 {
			continue
		}
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		pattern := []rune(matches[3])
		data := string(matches[4])
		x := 0
		for _, c := range data {
			if c == pattern[0] {
				x++
			}
		}
		if x > max || x < min {
			continue
		}
		valid++
	}

	log.Println(valid, "valid")
	return valid
}

//algo

func FindPrime(n int) bool{

for i := 2; i < n ; i++{
 if n % i == 0 {
  return false
}
}
return true
}

// func bubble(seq []int) []int{
	
// }