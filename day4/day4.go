package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func split_passport_string(passport_string string) map[string]string {
	var return_map = make(map[string]string)
	var keys_values_slice []string = strings.Fields(passport_string)
	for i := 0; i < len(keys_values_slice); i++ {
		var key_value_slice []string = strings.Split(keys_values_slice[i], ":")
		return_map[key_value_slice[0]] = key_value_slice[1]
	}
	return return_map
}

func get_dicts_of_passports(file_path string) []map[string]string {
	var slice_of_passport_maps []map[string]string
	var passport_batch []string = get_input_slice(file_path)

	passport_index := 0
	var passport_slice []string

	for i := 0; i < len(passport_batch); i++ {
		var line string = passport_batch[i]
		if line == "" {
			slice_of_passport_maps = append(slice_of_passport_maps, split_passport_string(passport_slice[passport_index]))
			passport_index++
		} else if len(passport_slice)-1 < passport_index {
			passport_slice = append(passport_slice, line)
		} else {
			passport_slice[passport_index] += fmt.Sprint(" ", line)
		}
	}
	slice_of_passport_maps = append(slice_of_passport_maps, split_passport_string(passport_slice[passport_index]))
	return slice_of_passport_maps
}

func regex_check(pattern string, input_string string) bool {
	r, _ := regexp.Compile(pattern)
	return r.MatchString(input_string)
}

func validate(key string, value string) bool {
	if value != "" {
		switch key {
		case "byr":
			return regex_check("^19[2-9][0-9]$|^200[0-2]$", value)
		case "iyr":
			return regex_check("^201[0-9]$|^2020$", value)
		case "eyr":
			return regex_check("^202[0-9]$|^2030$", value)
		case "hgt":
			return regex_check("^1[5-8][0-9]cm$|^19[0-3]cm$|^59in$|^6[0-9]in$|^7[0-6]in$", value)
		case "hcl":
			return regex_check("^#[0-9a-f]{6}$", value)
		case "ecl":
			return regex_check("^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$", value)
		case "pid":
			return regex_check("^[0-9]{9}$", value)
		default:
			return false
		}
	} else {
		return false
	}
}

func find_valid(passport_slice []map[string]string) int {
	count := 0
	for i := 0; i < len(passport_slice); i++ {
		min := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		inner_count := 0
		for ind := 0; ind < len(min); ind++ {
			if validate(min[ind], passport_slice[i][min[ind]]) {
				inner_count++
			} else {
				break
			}
		}
		if inner_count == len(min) {
			count += 1
		}
	}
	return count
}

func main() {
	// var passport_slice []map[string]string = get_dicts_of_passports("day4_input.txt")
	var passport_slice []map[string]string = get_dicts_of_passports("day4_input.txt")
	print(find_valid(passport_slice))

}
