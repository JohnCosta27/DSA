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
	list1 := []int{10, 12, 7, 3, 15, 9, 11}
  outputList := []int{10, 7, 3, 9, 11, 12, 15}
	Partition(list1, 1, len(list1) - 1)

  if !CompareSlices(list1, outputList) {
    t.Logf("Slices are not equal, got: %v\n", list1)
    t.Logf("Wanted: %v\n", outputList)
    t.FailNow()
  }
}

func TestMergeSort(t *testing.T) {
	list := []int{1, 18, -3, 7, 9, 4, -10, 20}
	sortedList := MergeSort(list)
	for i := 0; i < len(sortedList)-1; i++ {
		if sortedList[i] > sortedList[i+1] {
			t.Logf("Listed should be sorted, got: %x\n", sortedList)
			t.FailNow()
		}
	}
}

func TestLargeMergeSort(t *testing.T) {
	LENGTH := 100000
	list := make([]int, LENGTH)
	for i := 0; i < LENGTH; i++ {
		list[i] = rand.Intn(50000) - 25000
	}

	start := time.Now()
	sortedList := MergeSort(list)
	end := time.Since(start)
	t.Logf("Time taken: %v\n", end)

	if len(sortedList) != LENGTH {
		t.Logf("Length of list should be the same as unsorted, length: %d\n", len(sortedList))
		t.FailNow()
	}

	for i := 0; i < LENGTH-1; i++ {
		if sortedList[i] > sortedList[i+1] {
			t.Logf("Listed should be sorted, got: %x\n", sortedList)
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
