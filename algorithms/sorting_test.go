package algorithms

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeLists(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{6, 7, 8, 9, 10}
	expentedList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sortedList := mergeLists(list1, list2)

	for i := 0; i < len(sortedList); i++ {
		if sortedList[i] != expentedList[i] {
			t.Log("Expected all elements to be sorted")
			t.FailNow()
		}
	}
}

func TestPartition(t *testing.T) {
	list1 := []int{16, 12, 7, 3, 15, 9, 11}
	pivot := list1[len(list1)-1]
	Partition(list1, 0, len(list1)-1)

	index := -1
	for i, v := range list1 {
		if v == pivot {
			index = i
		}
	}

	for i, v := range list1 {
		if i < index {
			if !(v < list1[index]) {
				t.Logf("Every element to the left of the pivot must be smaller than it %v\n", list1)
				t.FailNow()
			}
		} else if i > index {
			if !(v > list1[index]) {
				t.Logf("Every element to the left of the pivot must be smaller than it %v\n", list1)
				t.FailNow()
			}
		}
	}

}

func TestSort(t *testing.T) {
	list := []int{1, 18, -3, 7, 9, 4, -10, 20}
	sortedListMerge := MergeSort(list)
	sortedListQuick := QuickSort(list)

	for i := 0; i < len(sortedListMerge)-1; i++ {
		if sortedListMerge[i] > sortedListMerge[i+1] {
			t.Logf("Listed should be sorted (merge sort), got: %v\n", sortedListMerge)
			t.FailNow()
		}
	}

	for i := 0; i < len(sortedListQuick)-1; i++ {
		if sortedListQuick[i] > sortedListQuick[i+1] {
			t.Logf("Listed should be sorted (quick sort), got: %v\n", sortedListQuick)
			t.FailNow()
		}
	}

}

func TestLargeSort(t *testing.T) {
	LENGTH := 1000000
	list := make([]int, LENGTH)
	for i := 0; i < LENGTH; i++ {
		list[i] = rand.Intn(50000) - 25000
	}

	start := time.Now()
	sortedListMerge := MergeSort(list)
	end := time.Since(start)
	t.Logf("Time taken (merge): %v\n", end)

	start = time.Now()
	sortedListQuick := QuickSort(list)
	end = time.Since(start)
	t.Logf("Time taken (quick): %v\n", end)

	if len(sortedListMerge) != LENGTH {
		t.Logf("Length of list should be the same as unsorted, length: %d\n", len(sortedListMerge))
		t.FailNow()
	}

	for i := 0; i < LENGTH-1; i++ {
		if sortedListMerge[i] > sortedListMerge[i+1] {
			t.Logf("Listed should be sorted, got: %x\n", sortedListMerge)
			t.FailNow()
		}
	}

	for i := 0; i < LENGTH-1; i++ {
		if sortedListQuick[i] > sortedListQuick[i+1] {
			t.Logf("Listed should be sorted, got: %v\n", sortedListQuick)
			t.FailNow()
		}
	}

}

func CompareSlices(arr1 []int, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
