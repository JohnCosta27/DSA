package algorithms

// The median algorithm consists of 2 parts.
// 1. rank(a, i) => returns the position of i in the array a, in ascending order
// 2. select(a, i) => returns the ith value from an array in ascending order.
// This is for an unsorted array.

// The crude solution for this problem is to simply sort the entire array, and return
// Running time for rank: O(nlogn) for sorting and then O(n) for traversing the array.
// Running time for select: O(nlogn) for sorting and then O(1) for returning.
func CrudeRank(arr []int, element int) int {
  sortedArr := MergeSort(arr)
  for i, el := range sortedArr {
    if el == element {
      return i
    }
  }
  return -1
}

func CrudeSelect(arr []int, index int) int {
  if (index >= len(arr)) {
    return -1
  }
  sortedArr := MergeSort(arr)
  return sortedArr[index]
}
