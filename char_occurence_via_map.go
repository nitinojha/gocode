// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	//Count of character occurence in a string
	var test_string = "aabc" //ojtjo
	var test_string1 = "bbcx"
	if len(test_string) == len(test_string1) {
		my_map3 := make(map[rune]int)
		//sum := 0
		for i := 0; i < len(test_string); i++ {
			if test_string[i] == test_string1[i] {
				continue
			} else {
				key := rune(test_string[i])
				key1 := rune(test_string1[i])
				my_map3[key] += 1
				if my_map3[key] == 0 {
					delete(my_map3, key)
				}
				my_map3[key1] -= 1
				if my_map3[key1] == 0 {
					delete(my_map3, key1)
				}
			}
		}
		if len(my_map3) == 0 {
			fmt.Println("Given strings are anagram")
		} else {
			fmt.Println("Given strings are mot anagram")
		}
		fmt.Println(my_map3)
	} else {
		fmt.Println("Given strings are not anagram")
	}
	my_map := make(map[rune]int)
	for _, v := range test_string {
		my_map[v] = my_map[v] + 1
	}
	fmt.Println(my_map)
	my_map1 := ISASCIIToString(my_map, &test_string)
	fmt.Println(my_map1, test_string)
	var my_map2 map[string]string = map[string]string{test_string: test_string}
	fmt.Println(my_map2)
}
func ISASCIIToString(my_map map[rune]int, input_str *string) map[string]int {
	temp_map := make(map[string]int)
	//var changed_str string
	//changed_str = *input_str
	for k, v := range my_map {
		temp_map[string(k)] = v
		if v > 1 {
			//changed_str = string(k + 1)
			fmt.Println(string(k), string(k+1))
			*input_str = strings.Replace(*input_str, string(k), string(k+1), v)
		}
	}
	fmt.Println("hiii", input_str)
	return temp_map

}
