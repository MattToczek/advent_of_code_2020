def get_input_as_list(file_path):
    input = []
    with open(file_path, 'r') as file:
        list = file.readlines()

    for item in list:
        input.append(item.replace('\n', ''))
    return input


def check_for_2020(int_list):
    list_max_index = len(int_list)-1
    for index in range(0, list_max_index):
        for remaining in range(index, list_max_index):
            if int(int_list[index]) + int(int_list[remaining]) == 2020:
                return int(int_list[index]) * int(int_list[remaining])


def check_for_less_than_2020(int_list):
    list = []
    list_max_index = len(int_list)-1
    for index in range(0, list_max_index):
        for remaining in range(index, list_max_index):
            if int(int_list[index]) + int(int_list[remaining]) < 2020:
                list.append([int(int_list[index]), int(int_list[remaining])])
    return list

def find_third_number(list_of_numbers, int_list):
    return_list = []
    list_max_index = len(int_list)-1
    for list in list_of_numbers:
        for index in range(0, list_max_index):
            if (int(int_list[index]) not in list_of_numbers) and (list[0] + list[1] + int(int_list[index]) == 2020):
                return  [int(int_list[index])] + list

    return return_list

def prod_of_list(list):
    result = 1
    for i in list:
        result = i*result
    return result

list_input = get_input_as_list("day1_input.txt")
print(check_for_2020(list_input))
list_to_prod = find_third_number(check_for_less_than_2020(list_input), list_input)
print(prod_of_list(list_to_prod))


