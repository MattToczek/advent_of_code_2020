def get_input_as_list(file_path):
    input = []
    with open(file_path, 'r') as file:
        list = file.readlines()

    for item in list:
        item_dict = {}
        item = item.rstrip('\n')
        list_tmp = item.split(' ')
        item_dict['min'] = int(list_tmp[0].split('-')[0])
        item_dict['max'] = int(list_tmp[0].split('-')[1])
        item_dict['char'] = list_tmp[1].rstrip(':')
        item_dict['pass'] = list_tmp[2]
        input.append(item_dict)
    return input

def check_if_valid(password_object_list):
    count = 0
    for password_object in password_object_list:
        number_of_occurences = password_object.get('pass').count(password_object.get('char'))
        if number_of_occurences >= password_object.get('min') and number_of_occurences <= password_object.get('max'):
            count+=1
    return count

def check_next_validation(password_object_list):
    return_count = 0
    for password_object in password_object_list:
        pass_array = list(password_object.get('pass'))
        if (pass_array[password_object['min']-1] == password_object['char']) ^ \
                (pass_array[password_object['max']-1] == password_object['char']):
            return_count += 1
    return return_count


password_object_list = get_input_as_list('day2_input.txt')
print(check_if_valid(password_object_list))
print(check_next_validation(password_object_list))
