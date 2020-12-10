package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func get_input_slice(file_path string) []int {
	var input []string
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var int_input = []int{}

	for _, i := range input {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		int_input = append(int_input, j)
	}

	return int_input
}

func contains(number int, number_slice []int) bool {
	for _, n := range number_slice {
		if number == n {
			return true
		}
	}
	return false
}

func find_sum(subset []int, total int) bool {
	for i := range subset {
		diff := total - subset[i]
		if contains(diff, subset) {
			return true
		}
	}
	return false
}

func loop_through_preamble(input []int, preamble_length int) int {
	for i, x := preamble_length, 0; i < len(input); {
		subset := input[x : x+preamble_length]
		total := input[i]

		if !find_sum(subset, total) {
			return total
		}

		x++
		i++
	}
	panic("not found")
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func find_slice_that_makes_target(input []int, target int) int {
	for i := 0; i < len(input); i++ {
		start := i
		for end := start + 1; end < len(input); end++ {
			if sum(input[start:end]) == target {
				subset_sum := input[start:end]
				sort.SliceStable(subset_sum, func(i, j int) bool {
					return subset_sum[i] < subset_sum[j]
				})
				return subset_sum[0] + subset_sum[len(subset_sum)-1]
			}
		}
	}
	panic("not found")
}

func main() {
	input := get_input_slice("day9_input.txt")
	fmt.Println(loop_through_preamble(input, 25))
	fmt.Println(find_slice_that_makes_target(input, 3199139634))
}
