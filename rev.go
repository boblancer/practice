package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
This algorithm uses two pointers to swap values between the edge of array 
Example input
["1", "2", "hello", "3", "4"]

step 1
["4", "2", "hello", "3", "1"]
  p              		  p
step 2
["4", "3", "hello", "2", "1"]
	   p    		 p
*/
func reverse(arr []string) {
	size := len(arr)
	half_size := int(math.Floor((float64(size) / 2.0)))

	for i := 0; i < half_size; i++ {
		front_p := &arr[i]
		back_p := &arr[size - 1 - i]
		*front_p, *back_p = *back_p, *front_p
	}
}

func test_odd_input_size() bool {
	in := 		[]string{ "bar", "abc", "def", "xyz", "bird"}
	expected := []string{ "bird", "xyz", "def", "abc", "bar"}
	reverse(in)

	return reflect.DeepEqual(in, expected)
}

func test_even_input_size() bool{
	in := 		[]string{ "bar", "abc", "xyz", "bird"}
	expected := []string{ "bird", "xyz", "abc", "bar"}
	reverse(in)

	return reflect.DeepEqual(in, expected)
}

func main() {

	fmt.Println("Case 1", test_even_input_size())
	fmt.Println("Case 2", test_odd_input_size())
}
	