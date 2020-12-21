package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func get_map_slice(input []string) []struct {
	string
	int
} {
	var map_slice []struct {
		string
		int
	}
	for i := 0; i < len(input); i++ {
		val, _ := strconv.Atoi(input[i][1:])
		map_slice = append(map_slice, struct {
			string
			int
		}{input[i][:1], val})
	}
	return map_slice
}

func follow_directions(input []struct {
	string
	int
}) int {
	var position Position
	position.S = 0
	position.E = 0
	position.N = 0
	position.W = 0
	facing_list := []string{"E", "S", "W", "N"}
	facing_index := 0
	facing := "E"
	for i := 0; i < len(input); i++ {
		fmt.Println("Input:", input[i])
		switch input[i].string {
		case "R":
			change := input[i].int / 90
			if facing_index+change > 3 {
				facing_index -= 4
			}
			facing_index += change
			facing = facing_list[facing_index]
			break
		case "L":
			change := input[i].int / 90
			if facing_index-change < 0 {
				facing_index += 4
			}
			facing_index -= change
			facing = facing_list[facing_index]
			break
		case "F":
			position = move(facing, input[i].int, position)
			break

		}
		position = move(input[i].string, input[i].int, position)
	}

	var sum_a int
	var sum_b int

	if position.S >= 0 {
		sum_a = position.S
	} else {
		sum_a = position.N
	}
	if position.E >= 0 {
		sum_b = position.E
	} else {
		sum_b = position.W
	}

	return sum_a + sum_b
}

func move(compass_point string, distance int, current_position Position) Position {
	switch compass_point {
	case "N":
		current_position.N += distance
		current_position.S -= distance
		break
	case "S":
		current_position.S += distance
		current_position.N -= distance
		break
	case "E":
		current_position.E += distance
		current_position.W -= distance
		break
	case "W":
		current_position.W += distance
		current_position.E -= distance
		break
	}
	fmt.Println("Direction:", compass_point, "Distance:", distance, "New Position:", current_position)
	return current_position
}

func main() {
	input := get_map_slice(get_input_slice("day12_input.txt"))
	fmt.Println(follow_directions(input))
}

type Position struct {
	N int
	S int
	E int
	W int
}
