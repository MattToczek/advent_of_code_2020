package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func check_and_split(input_suffix string, return_slice []string) []string {
	var no_number_slice []string
	matched, _ := regexp.MatchString(`bag`, input_suffix)
	if matched {
		number, _ := strconv.Atoi(string(input_suffix[0]))
		splitter := regexp.MustCompile(`bags?`)
		no_number_slice = splitter.Split(input_suffix[1:len(input_suffix)], 2)
		bag := no_number_slice[0]
		for i := 0; i < number; i++ {
			return_slice = append(return_slice, bag)
		}
		return_slice = check_and_split(strings.Join(no_number_slice[1:len(no_number_slice)], ""), return_slice)
	}

	return return_slice
}

func split_to_bagmap(input_slice []string) map[string][]string {
	var bagmap_map_slice = make(map[string][]string)
	for line_no := range input_slice {
		line := input_slice[line_no]
		re := regexp.MustCompile(`\.|,| `)
		splitter := regexp.MustCompile(`bagscontain`)
		line = re.ReplaceAllString(line, "")
		slice := splitter.Split(line, -1)

		var return_slice []string
		bagmap_map_slice[slice[0]] = check_and_split(slice[1], return_slice)
	}
	return bagmap_map_slice
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func get_containing_bags(bag_map_slice map[string][]string, bag string, containers []string) []string {
	for key, value := range bag_map_slice {
		if contains(value, bag) {
			containers = append(containers, key)
			containers = get_containing_bags(bag_map_slice, key, containers)
		}
	}
	return containers
}

func get_bags_inside(bag_map_slice map[string][]string, bag string, contained []string) []string {
	for i := 0; i < len(bag_map_slice[bag]); i++ {
		contained = append(contained, bag_map_slice[bag][i])
		contained = get_bags_inside(bag_map_slice, bag_map_slice[bag][i], contained)
	}
	return contained
}

func main() {
	var input []string = get_input_slice("day7_input.txt")
	var containers []string
	bag_map_slice := split_to_bagmap(input)
	new_containers := get_containing_bags(bag_map_slice, "shinygold", containers)
	fmt.Println(len(unique(new_containers)))
	var contains []string
	fmt.Println(len(get_bags_inside(bag_map_slice, "shinygold", contains)))

}
