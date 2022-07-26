package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{1, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum  empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 1, 2}, []int{3, 4, 5})
		want := []int{4, 5, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
