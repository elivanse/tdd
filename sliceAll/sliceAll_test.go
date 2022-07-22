package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	//	if got != want {
	// go no puede comparar valores de
	// igualdad con slices

	//con reflect.DeepEqual compara la
	//totalidad de sus elementos
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}