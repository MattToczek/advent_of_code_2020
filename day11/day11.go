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

func get_input_int_byte_map(file_path string) map[int][]byte {
	input_slice := get_input_slice(file_path)
	var input_int_byte_map map[int][]byte = make(map[int][]byte)
	for i := 0; i < len(input_slice); i++ {
		input_int_byte_map[i] = []byte(input_slice[i])
	}
	return input_int_byte_map
}

func pretty_print(result map[int][]byte) string {
	str := ""
	for i := 0; i < len(result); i++ {
		str += fmt.Sprintln(string(result[i]))
	}
	return str
}

func find_seat_pattern(input map[int][]byte) map[int][]byte {
	changed_plan := make(map[int][]byte)
	for i := 0; i < len(input); i++ {
		for inner_i := 0; inner_i < len(input[i]); inner_i++ {
			if input[i][inner_i] == '.' {
				changed_plan[i] = append(changed_plan[i], '.')
			} else {
				count := 0
				if i > 0 {
					if inner_i > 0 && input[i-1][inner_i-1] == '#' {
						count++
					}
					if input[i-1][inner_i] == '#' {
						count++
					}
					if inner_i < len(input[i])-1 && input[i-1][inner_i+1] == '#' {
						count++
					}
				}
				if inner_i > 0 && input[i][inner_i-1] == '#' {
					count++
				}
				if inner_i < len(input[i])-1 && input[i][inner_i+1] == '#' {
					count++
				}
				if i < len(input)-1 {
					if inner_i > 0 && input[i+1][inner_i-1] == '#' {
						count++
					}
					if input[i+1][inner_i] == '#' {
						count++
					}
					if inner_i < len(input[i])-1 && input[i+1][inner_i+1] == '#' {
						count++
					}
				}
				if count > 3 && input[i][inner_i] == '#' {
					changed_plan[i] = append(changed_plan[i], 'L')
				} else if count == 0 && input[i][inner_i] == 'L' {
					changed_plan[i] = append(changed_plan[i], '#')
				} else {
					changed_plan[i] = append(changed_plan[i], input[i][inner_i])
				}
			}
		}
	}
	pretty_print(changed_plan)
	return changed_plan
}

func check_for_consistency(input map[int][]byte) int {
	consistent := false
	var previous map[int][]byte = input
	for consistent == false {
		current := find_seat_pattern(previous)
		if pretty_print(current) == pretty_print(previous) {
			consistent = true
		}
		previous = current
	}
	return strings.Count(pretty_print(previous), "#")
}

func main() {
	var input map[int][]byte = get_input_int_byte_map("day11_input.txt")
	pretty_print(input)
	input2 := find_seat_pattern(input)
	pretty_print(input2)
	input3 := find_seat_pattern(input2)
	pretty_print(input3)
	print(check_for_consistency(input))
}

// ================================================================================
// ================================= PART 2 =======================================
// ================================================================================
// func find_seat_pattern(input map[int][]byte, radius int) (map[int][]byte, int) {
// 	edge_count := 0
// 	changed_plan := make(map[int][]byte)
// 	for i := 0; i < len(input); i++ {
// 		for inner_i := 0; inner_i < len(input[i]); inner_i++ {
// 			if input[i][inner_i] == '.' {
// 				changed_plan[i] = append(changed_plan[i], '.')
// 			} else {
// 				count := 0
// 				if i >= radius && inner_i >= radius {
// 					if input[i-radius][inner_i-radius] == '#' {
// 						count++
// 					}
// 				} else {
// 					edge_count++
// 				}
// 				if i >= radius {
// 					if input[i-radius][inner_i] == '#' {
// 						count++
// 					}
// 				} else {
// 					edge_count++
// 				}
// 				if i >= radius && inner_i < len(input[i])-radius {
// 					if input[i-radius][inner_i+radius] == '#' {
// 						count++
// 					}
// 				} else {
// 					edge_count++
// 				}
// 				if inner_i >= radius && input[i][inner_i-radius] == '#' {
// 					count++
// 				} else {
// 					edge_count++
// 				}
// 				if inner_i < len(input[i])-radius && input[i][inner_i+radius] == '#' {
// 					count++
// 				} else {
// 					edge_count++
// 				}
// 				if i < len(input)-radius && inner_i >= radius {
// 					if input[i+radius][inner_i-radius] == '#' {
// 						count++
// 					}
// 				} else {
// 					edge_count++
// 				}
// 				if i < len(input)-radius {
// 					if input[i+radius][inner_i] == '#' {
// 						count++
// 					}
// 				} else {
// 					edge_count++
// 				}
// 				if i < len(input)-radius && inner_i < len(input[i])-radius {
// 					if input[i+radius][inner_i+radius] == '#' {
// 						count++
// 					}
// 				} else {
// 					edge_count++
// 				}

// 				if count > 3 && input[i][inner_i] == '#' {
// 					changed_plan[i] = append(changed_plan[i], 'L')
// 				} else if count == 0 && input[i][inner_i] == 'L' {
// 					changed_plan[i] = append(changed_plan[i], '#')
// 				} else {
// 					changed_plan[i] = append(changed_plan[i], input[i][inner_i])
// 				}
// 			}
// 		}
// 	}
// 	pretty_print(changed_plan)
// 	return changed_plan, edge_count
// }

// func check_for_consistency_pt2(input map[int][]byte) int {
// 	consistent := false
// 	var previous map[int][]byte = input
// 	current := previous
// 	for consistent == false {
// 		current_radius := 1
// 		edge_count := 0
// 		for edge_count > 8 {
// 			current, edge_count = find_seat_pattern(previous, current_radius)
// 			current_radius++
// 		}
// 		if pretty_print(current) == pretty_print(previous) {
// 			consistent = true
// 		}
// 		previous = current
// 	}
// 	return strings.Count(pretty_print(previous), "#")
// }
