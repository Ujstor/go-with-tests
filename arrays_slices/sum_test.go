package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	
	t.Run("collection of any size", func (t *testing.T) {
		numbers := []int{1, 2, 3}
	
		got := Sum(numbers)
		want := 6
	
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}