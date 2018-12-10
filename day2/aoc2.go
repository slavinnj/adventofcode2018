package main

import (
    "fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var file = "box_ids.txt"

func main() {
	data, err := ioutil.ReadFile(file)
	check(err)
	ids := strings.Split(string(data), "\n")
	fmt.Printf("checksum = %d\n", chksum(ids))
}

func chksum(ids []string) int {
	two_count := 0;
	three_count := 0;
	for i := 0; i < len(ids); i++ {
		this_box := ids[i]
		letter_frequencies := make(map[string]int)
		for _, c := range this_box {
			this_char := string(c)
			if freq, ok := letter_frequencies[this_char]; ok {
				letter_frequencies[this_char] = freq + 1
			} else {
				letter_frequencies[this_char] = 1
			}
		}
		seen_two := false
		seen_three := false
		for _, v := range letter_frequencies {
			if v == 2 && seen_two == false {
				two_count++
				seen_two = true
			} else if v == 3 && seen_three == false {
				three_count++
				seen_three = true
			}
		}
	}
	return two_count * three_count
}