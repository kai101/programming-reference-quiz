package main

import (
	"testing"
	"reflect"
)

func TestSelectionSort(t *testing.T){
	cases := []struct {
		in, want []int
	}{
		{[]int{5,4,3,2,1}, []int{1,2,3,4,5}},
		{[]int{5001,4001,3001,3001,1001}, []int{1001,3001,3001,4001,5001}},
	}
	
	for _, c := range cases {
		got := SelectionSort(c.in)
		if !reflect.DeepEqual(got, c.want){
			t.Errorf("BucketSort(%v) == %v, want %v", c.in, got, c.want)
		}
	}

}
