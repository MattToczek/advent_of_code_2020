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

func get_group_map(input []string) []Group {
	index := 0
	var groups []Group
	for line_no := range input {
		line := input[line_no]
		if line == "" {
			index += 1
		} else {
			string_slice := strings.Split(line, "")
			if len(groups) == index {
				var group Group
				group.answers = make(map[string]int)
				for char := range string_slice {
					group.answers[string_slice[char]] += 1
				}
				group.members += 1
				groups = append(groups, group)

			} else {
				for char := range string_slice {
					groups[index].answers[string_slice[char]] += 1
				}
				groups[index].members += 1
			}
		}
	}
	return groups
}

func get_sum(groups []Group) int {
	sum := 0
	for val := range groups {
		sum += len(groups[val].answers)
	}
	return sum
}

func get_all_answered(groups []Group) int {
	sum := 0
	for val := range groups {
		for _, number := range groups[val].answers {
			if number == groups[val].members {
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	var input []string = get_input_slice("day6_input.txt")
	var groups []Group = get_group_map(input)
	var sum int = get_sum(groups)
	var all_ans = get_all_answered(groups)
	fmt.Println(sum)
	fmt.Println(all_ans)

}

type Group struct {
	answers map[string]int
	members int
}
