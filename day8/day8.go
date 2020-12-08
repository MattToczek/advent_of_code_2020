package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func split_to_map(input_slice []string) map[int]action_object {
	splitter := regexp.MustCompile(` +?`)
	return_map := make(map[int]action_object)
	for i := 0; i < len(input_slice); i++ {
		pair := splitter.Split(input_slice[i], -1)
		var ao action_object
		ao.action = pair[0]
		ao.number, _ = strconv.Atoi(pair[1])
		return_map[i] = ao
	}
	return return_map
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func do_action(action_object_map map[int]action_object, key int, ind_so_far []int, count int) int {
	if contains(ind_so_far, key) {
		return count
	}
	var next_key int
	ind_so_far = append(ind_so_far, key)
	switch action_object_map[key].action {
	case "jmp":
		next_key = key + action_object_map[key].number
	case "nop":
		next_key = key + 1
	case "acc":
		next_key = key + 1
		count += action_object_map[key].number
	}
	return do_action(action_object_map, next_key, ind_so_far, count)
}

func find_loop(action_object_map map[int]action_object) int {
	var ind_so_far []int
	key := 0
	count := 0

	return do_action(action_object_map, key, ind_so_far, count)
}

func main() {
	input := get_input_slice("day8_input.txt")
	action_object_map := split_to_map(input)
	fmt.Print(find_loop(action_object_map))
}

type action_object struct {
	number int
	action string
}
