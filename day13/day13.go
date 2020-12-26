package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

func split_input_to_parts(input []string) (int, []int) {
	s := strings.Split(input[1], ",")
	var bus_times []int
	earliest, _ := strconv.Atoi(input[0])
	for i := 0; i < len(s); i++ {
		val, _ := strconv.Atoi(s[i])
		if val != 0 {
			bus_times = append(bus_times, val)
		}
	}
	return earliest, bus_times
}

func find_earliest_bus(earliest int, bus_times []int) int {
	closest_time := bus_times[0] * int(math.Ceil(float64(earliest)/float64(bus_times[0])))
	closest_bus := bus_times[0]
	for i := 0; i < len(bus_times); i++ {
		next_multiple := bus_times[i] * int(math.Ceil(float64(earliest)/float64(bus_times[i])))
		if closest_time > next_multiple {
			closest_time = next_multiple
			closest_bus = bus_times[i]
		}
	}
	return (closest_time - earliest) * closest_bus
}

func main() {
	input := get_input_slice("input.txt")
	earliest, bus_times := split_input_to_parts(input)
	fmt.Println(find_earliest_bus(earliest, bus_times))

}
