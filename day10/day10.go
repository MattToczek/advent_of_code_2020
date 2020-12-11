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

func sort_slice(input []int) []int {
	sort.SliceStable(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)
	return input
}

func find_jumps(input []int) int {
	input = sort_slice(input)
	var count_map = map[int]int{3: 0, 2: 0, 1: 0}
	for i := 1; i < len(input); i++ {
		count_map[input[i]-input[i-1]] += 1
	}
	return count_map[1] * count_map[3]
}

func find_all_combinations(input []int) {
	var count_map map[int]int = make(map[int]int)
	input = sort_slice(input)
	for i := 0; i < len(input)-3; i++ {
		count_map[i] = 1
		if input[i+2]-input[i] <= 3 {
			count_map[i]++
		}
		if input[i+3]-input[i] <= 3 {
			count_map[i]++
		}
	}
	count := 0
	count += loop(0, input, count_map, make(map[int]int))

	fmt.Println(count)
}

func loop(number int, slice []int, count_map map[int]int, visited map[int]int) int {

	count := 1

	if val, ok := visited[slice[number]]; ok {
		return val
	}

	for i := number; i < len(slice)-3; i++ {
		if count_map[i] > 1 {
			count += loop(i+2, slice, count_map, visited)
		}
		if count_map[i] > 2 {
			count += loop(i+3, slice, count_map, visited)
		}
	}

	visited[slice[number]] = count

	return count
}

func main() {
	input := get_input_slice("day10_input.txt")
	fmt.Println(find_jumps(input))
	find_all_combinations(input)
}
