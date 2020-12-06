package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func find_binary(slice []string, full_number int) int {
	var group = [2]float64{0, float64(full_number)}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "L" || slice[i] == "F" {
			group[1] = math.Floor(group[0] + (group[1]-group[0])/2)
		} else if slice[i] == "R" || slice[i] == "B" {
			group[0] = math.Ceil(group[1] - (group[1]-group[0])/2)
		} else {
			fmt.Println("Char not recognised:", slice[i], "in", slice)
		}
	}
	if slice[len(slice)-1] == "L" || slice[len(slice)-1] == "F" {
		return int(group[0])
	}
	return int(group[1])
}

func create_boarding_pass(binary_string string) BoardingPass {
	bp_slice := strings.Split(binary_string, "")

	var bp BoardingPass
	bp.row_slice = bp_slice[0:7]
	bp.column_slice = bp_slice[7:10]
	bp.row = find_binary(bp.row_slice, 127)
	bp.column = find_binary(bp.column_slice, 7)
	bp.id = bp.row*8 + bp.column

	return bp
}

func find_highest_id(boarding_pass_slice []BoardingPass) BoardingPass {
	var highest_id BoardingPass
	highest_id.id = 0

	for i := 0; i < len(boarding_pass_slice); i++ {
		var contender BoardingPass = boarding_pass_slice[i]
		if contender.id > highest_id.id {
			highest_id = contender
		}
	}

	return highest_id
}

func get_boarding_pass_slice(input []string) []BoardingPass {
	var boarding_pass_slice []BoardingPass
	for i := 0; i < len(input); i++ {
		line := input[i]
		boarding_pass_slice = append(boarding_pass_slice, create_boarding_pass(line))
	}
	return boarding_pass_slice
}

func find_seat_id(boarding_pass_slice []BoardingPass) int {
	map_of_seats := make(map[int][]int)
	for i := 0; i < len(boarding_pass_slice); i++ {
		map_of_seats[boarding_pass_slice[i].row] = append(map_of_seats[boarding_pass_slice[i].row], boarding_pass_slice[i].column)
	}
	for key, element := range map_of_seats {
		if len(element) < 8 && len(map_of_seats[(key-1)]) == 8 && len(map_of_seats[(key+1)]) == 8 {
			sort.Ints(map_of_seats[key])
			for i := 0; i < 8; i++ {
				if element[i] != i {
					return key*8 + i
				}
			}
		}
	}
	panic("not found")
}

func main() {
	var input []string = get_input_slice("day5_input.txt")
	var boarding_pass_slice []BoardingPass = get_boarding_pass_slice(input)
	var seat_id = find_seat_id(boarding_pass_slice)
	fmt.Println(seat_id)
}

type BoardingPass struct {
	row_slice    []string
	column_slice []string
	row          int
	column       int
	id           int
}
