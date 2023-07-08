package main

import "testing"

type addTest struct {
    array, expected []int
    target int
}

var addTests = []addTest{
    {[]int{2,7,11,15}, []int{0,1}, 9},
    {[]int{3,3}, []int{0,1}, 6},
    {[]int{3,2,4}, []int{1,2}, 6},
    {[]int{-4,2,6,3,10,6}, []int{0,2}, 2},
}

func TestAdd(t *testing.T) {

    for _, test := range addTests {
        if output := twoSum(test.array, test.target); output[0] != test.expected[0] || output[1] != test.expected[1] {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
    }
}

