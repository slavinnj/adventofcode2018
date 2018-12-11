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

func difference(id1 string, id2 string) int {
	diff := 0
	for index, char := range id1 {
		if id2[index] != byte(char) {
			diff++
		}
	}
	return diff
}

func main() {
	data, err := ioutil.ReadFile(file)
	check(err)
	ids := strings.Split(string(data), "\n")
	chksum := chksum(ids)
	fmt.Printf("checksum = %d\n", chksum)

	match := false
	index := 0
	var box1 *string
	var box2 *string
	for !match {
		id1 := ids[index]
		for i := range ids {
			diff := difference(id1, ids[i])
			if diff == 1 {
				box1 = &id1
				box2 = &ids[i]
				match = true
				break
			}
		}
		if !match && ((index + 1) < len(ids)) {
			index++
			id1 = ids[index]
		}
	}
	if box1 != nil && box2 != nil {
		fmt.Printf("Box 1 = %s\n", *box1)
		fmt.Printf("Box 2 = %s\n", *box2)
		box2_id := *box2
		common := []string{}
		common_str := ""
		for pos, char := range *box1 {
			if char == rune(box2_id[pos]) {
				common = append(common, string(char))
				common_str = strings.Join(common, "")
			}
		}
		fmt.Printf("Common letters = %s\n", common_str)
	} else {
		fmt.Println("No matching boxes.")
	}
}

func chksum(ids []string) int {
	two_count := 0
	three_count := 0
	two_letter_boxes := make(map[string]struct{})
	three_letter_boxes := make(map[string]struct{})
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
				two_letter_boxes[this_box] = struct{}{}
			} else if v == 3 && seen_three == false {
				three_count++
				seen_three = true
				three_letter_boxes[this_box] = struct{}{}
			}
		}
	}
	return two_count * three_count
}
