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
	return get_manhattan_distance(position)
}

func get_manhattan_distance(position Position) int {
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
	return current_position
}

func change_axis(current_axis Position, turns int) Position {
	values := []int{current_axis.N, current_axis.E, current_axis.S, current_axis.W}
	values = append(values[turns:], values[:turns]...)
	return Position{values[0], values[1], values[2], values[3]}
}

func way_point_directions(input []struct {
	string
	int
}) int {
	way_point := Position{1, 10, -1, -10}
	current_position := Position{0, 0, 0, 0}
	for i := 0; i < len(input); i++ {
		switch input[i].string {
		case "R":
			change := input[i].int / 90
			way_point = change_axis(way_point, 4-change)
			break
		case "L":
			change := input[i].int / 90
			way_point = change_axis(way_point, change)
			break
		case "F":
			if way_point.N >= 0 {
				current_position.N += (way_point.N * input[i].int)
				current_position.S -= (way_point.N * input[i].int)
			}
			if way_point.S >= 0 {
				current_position.S += (way_point.S * input[i].int)
				current_position.N -= (way_point.S * input[i].int)
			}
			if way_point.E >= 0 {
				current_position.E += (way_point.E * input[i].int)
				current_position.W -= (way_point.E * input[i].int)
			}
			if way_point.W >= 0 {
				current_position.W += (way_point.W * input[i].int)
				current_position.E -= (way_point.W * input[i].int)
			}
			break
		default:
			way_point = move(input[i].string, input[i].int, way_point)
			break
		}

	}

	return get_manhattan_distance(current_position)
}

func main() {
	input := get_map_slice(get_input_slice("day12_input.txt"))
	fmt.Println(follow_directions(input))
	fmt.Println(way_point_directions(input))
}

type Position struct {
	N int
	E int
	S int
	W int
}
