package hw1

import (
	"fmt"
)

var flag = true

// func TestCompare1(t *testing.T) {
// 	testCases := []struct {
// 		slice1 []int
// 		slice2 []int
// 		equal  bool
// 	}{
// 		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
// 		{[]int{3, 2, 1}, []int{1, 2, 3}, true},
// 		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
// 		{[]int{1, 2, 3}, []int{1, 2}, false},
// 		{[]int{}, []int{}, true},
// 	}

// 	for _, testCase := range testCases {
// 		t.Run(fmt.Sprintf("Compare(%v, %v)", testCase.slice1, testCase.slice2), func(t *testing.T) {
// 			result := Compare(testCase.slice1, testCase.slice2)
// 			if result != testCase.equal {
// 				t.Errorf("Expected Compare(%v, %v) to be %v, but got %v", testCase.slice1, testCase.slice2, testCase.equal, result)
// 			}
// 		})
// 	}
// }

func TestCompare() {
	testCases := []struct {
		slice1 []int
		slice2 []int
		equal  bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{3, 2, 1}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2}, false},
	}
	for _, testCase := range testCases {
		result := Compare(testCase.slice1, testCase.slice2)
		if result != testCase.equal {
			fmt.Printf("Expected Compare(%v, %v) to be %v, but got %v", testCase.slice1, testCase.slice2, testCase.equal, result)
			flag = false
		}
	}
}

func TestTwoSum() {
	testCases := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 18, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
		{[]int{2, 7, 11, 15}, 13, []int{0, 2}},
		{[]int{2, 7, 11, 15}, 22, []int{1, 3}},
	}
	for _, testCase := range testCases {
		result := TwoSum(testCase.nums, testCase.target)
		if !Compare(result, testCase.result) {
			fmt.Printf("Expected TwoSum(%v, %v) to be %v, but got %v", testCase.nums, testCase.target, testCase.result, result)
			flag = false
		}
	}
}

func TestLongestCommonPrefix() {
	testCases := []struct {
		strs   []string
		result string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"a"}, "a"},
		{[]string{"ac", "c", "acc"}, ""},
	}
	for _, testCase := range testCases {
		result := LongestCommonPrefix(testCase.strs)
		if result != testCase.result {
			fmt.Printf("Expected LongestCommonPrefix(%v) to be %v, but got %v", testCase.strs, testCase.result, result)
			flag = false
		}
	}
}

func TestStruct() {
	testCases := []struct {
		myStruct MyStruct
		result   []int
	}{
		{MyStruct{slice1: []int{1, 2, 3}}, []int{1, 2, 3}},
		{MyStruct{slice1: []int{1, -2, 4}}, []int{-2, 1, 4}},
		{MyStruct{slice1: []int{1, -89}}, []int{-89, 1}},
		{MyStruct{slice1: []int{}}, []int{}},
	}
	for _, testCase := range testCases {
		realResult := testCase.myStruct.sort()
		if !compareforTest(testCase.result, realResult) {
			fmt.Printf("Expected sort() to be %v, but got %v", testCase.result, realResult)
			flag = false
		}
	}
}

func compareforTest(slice1 []int, slice2 []int) bool {
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func RunTests() {
	TestCompare()
	TestTwoSum()
	TestLongestCommonPrefix()
	TestStruct()
	if flag {
		fmt.Println("Testing Complete Successfully")
	}
}
