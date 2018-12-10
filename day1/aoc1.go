package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var file = "frequencies.txt"

func main() {
	data, err := ioutil.ReadFile(file)
	check(err)
	frequencies := strings.Split(string(data), "\n")
	fmt.Printf("Final frequency = %d\n", sum_frequencies(frequencies))
	fmt.Printf("First repeated frequency = %d\n", fst_repeated_frequency(frequencies))
}

func sum_frequencies(frequencies []string) int {
	sum_freq := 0
	for i := 0; i < len(frequencies); i++ {
		freq, err := strconv.ParseInt(frequencies[i], 10, 64)
		check(err)
		sum_freq += int(freq)
	}
	return sum_freq
}

func fst_repeated_frequency(frequencies []string) int {
	seen_frequencies := make(map[int]bool)
	sum_freq := 0
	seen := false
	for seen == false {
		for i := 0; i < len(frequencies); i++ {
			freq, err := strconv.ParseInt(frequencies[i], 10, 64)
			check(err)
			sum_freq += int(freq);
			if seen_frequencies[sum_freq] == true {
				seen = true
				break;
			} else {
				seen_frequencies[sum_freq] = true
			}
		}
	}
	return sum_freq
}