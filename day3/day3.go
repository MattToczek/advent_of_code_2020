package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func get_input_slice(file_path string) []string {
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
	return input
}

func get_number_of_trees(horizontal_dist int, vertical_dist int, slope_array []string) int {
	tree_count := 0
	index_count := 0
	for i := 0; i < len(slope_array); i += vertical_dist {
		row_array := strings.Split(slope_array[i], "")
		if row_array[index_count] == "#" {
			tree_count += 1
		}

		index_count += horizontal_dist
		if index_count >= len(slope_array[i]) {
			index_count -= len(slope_array[i])
		}
	}
	return tree_count
}

func main() {
	var slope_array []string = get_input_slice("day3_input.txt")

	slope1 := get_number_of_trees(1, 1, slope_array)
	slope2 := get_number_of_trees(3, 1, slope_array)
	slope3 := get_number_of_trees(5, 1, slope_array)
	slope4 := get_number_of_trees(7, 1, slope_array)
	slope5 := get_number_of_trees(1, 2, slope_array)

	println(slope1 * slope2 * slope3 * slope4 * slope5)

	fmt.Println(fmt.Sprintf("%s%d", "Slope 1 = ", slope1))
	fmt.Println(fmt.Sprintf("%s%d", "Slope 2 = ", slope2))
	fmt.Println(fmt.Sprintf("%s%d", "Slope 3 = ", slope3))
	fmt.Println(fmt.Sprintf("%s%d", "Slope 4 = ", slope4))
	fmt.Println(fmt.Sprintf("%s%d", "Slope 4 = ", slope5))
}
