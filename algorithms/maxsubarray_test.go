package algorithms

import (
  "testing"
)

func TestBestCrossing(t *testing.T) {
  arr := []int{5, 5, 5, 5, 5, 5, 5}
  bestCross := FindBestCrossing(len(arr) / 2, arr)

  if bestCross != 35 {
    t.Logf("Expected max to be 35, got: %d\n", bestCross)
  }

  newArr := []int{-5, 6, -1, 5, 2}
  bestCross = FindBestCrossing(len(newArr) / 2, newArr)

  if bestCross != 12 {
    t.Logf("Expected max to be 12, got: %d\n", bestCross)
  }
}

func TestMaxSubarray(t *testing.T) {
  arr := []int{10, 10, 10}
  bestSubArray := FindMaxSubArray(arr)
  
  if bestSubArray != 30 {
    t.Logf("Expected 30, but got: %d\n", bestSubArray)
  }

  arr = []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
  bestSubArray = FindMaxSubArray(arr)
  
  if bestSubArray != 43 {
    t.Logf("Expected 43, but got: %d\n", bestSubArray)
  }
}
