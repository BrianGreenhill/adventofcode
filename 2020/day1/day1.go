package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// read the datafile to bytes
	bytes, err := ioutil.ReadFile("day1/data")
	if err != nil {
		log.Fatal(err)
	}

	// convert the data to a slice of integer values
	nums, err := ReadInts(strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatal(err)
	}

	result := twoSum(nums, 2020)
	if result == nil {
		fmt.Println("There are no numbers that add up to 2020 in the dataset")
	} else {
		sum := nums[result[0]] * nums[result[1]]
		fmt.Printf("The two numbers (%d + %d = 2020) multiplied: %d\n", nums[result[0]], nums[result[1]], sum)
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}
