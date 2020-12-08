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
	fmt.Println(action_object_map[key])
	return do_action(action_object_map, next_key, ind_so_far, count)
}

func check_for_no_loop(action_object_map map[int]action_object, key int, ind_so_far []int, count int) bool {
	if key == len(action_object_map) {
		return true
	} else if contains(ind_so_far, key) {
		return false
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
	fmt.Println(action_object_map[key])
	return check_for_no_loop(action_object_map, next_key, ind_so_far, count)
}

func create_new_map_and_test(action_object_map map[int]action_object) int? {

}

func check_map_for_no_loop(action_object_map map[int]action_object) int {
	for key, value := range action_object_map {
		if value.action == "jmp" {
			new_map := make(map[int]action_object)
			for k, v := range action_object_map {
				new_map[k] = v
			}
			var new_ao action_object
			new_ao.action = "nop"
			new_ao.number = value.number
			new_map[key] = new_ao
			var empty_slice []int
			if check_for_no_loop(new_map, 0, empty_slice, 0) {
				var empty_slice []int
				return do_action(new_map, 0, empty_slice, 0)
			}
		} else if value.action == "nop" {
			new_map := make(map[int]action_object)
			for k, v := range action_object_map {
				new_map[k] = v
			}
			var new_ao action_object
			new_ao.action = "jmp"
			new_ao.number = value.number
			new_map[key] = new_ao
			var empty_slice []int
			if check_for_no_loop(new_map, 0, empty_slice, 0) {
				var empty_slice []int
				return do_action(new_map, 0, empty_slice, 0)
			}
		}
	}
	panic("Not found")
}

func change_action(action_object_map map[int]action_object, key int, action string) map[int]action_object {
	var ao action_object
	ao.action = action
	ao.number = key + 1
	action_object_map[key] = ao
	return action_object_map
}

func main() {
	input := get_input_slice("day8_input.txt")
	action_object_map := split_to_map(input)
	var empty_slice []int
	fmt.Print(do_action(action_object_map, 0, empty_slice, 0))

	// var test_ao action_object
	// test_ao.action = "nop"
	// test_ao.number = 0
	// var test_object_map map[int]action_object = make(map[int]action_object)
	// test_object_map[0] = test_ao
	fmt.Print(check_map_for_no_loop(action_object_map))

}

type action_object struct {
	number int
	action string
}
