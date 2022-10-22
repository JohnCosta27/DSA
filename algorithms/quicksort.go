package algorithms

func QuickSort(arr []int) []int {
  QuickSortRecursive(arr, 1, len(arr) - 1)
  return arr 
}

// @param start is the start of the bigger partition
func QuickSortRecursive(arr []int, start int, end int) {
  if start < end {
    mid := Partition(arr, start, end)
    QuickSortRecursive(arr, start, mid - 1)
    QuickSortRecursive(arr, mid + 1, end)
  }
}

// Modify the array in place, such that there are 2 partitions,
// one with elements that are smaller, and the other bigger.
func Partition(arr []int, start int, end int) int {
  pivot := arr[end]
  smallerPartition := start - 1
  for biggerPartition := start; biggerPartition < end; biggerPartition++ {
    // If the current element in the bigger partition space is smaller than pivot.
    // Exchange element in the bigger partition with the one in the smaller.
    // Increase size of smaller partition.
    if arr[biggerPartition] <= pivot {
      smallerPartition++
      temp := arr[smallerPartition]
      arr[smallerPartition] = arr[biggerPartition]
      arr[biggerPartition] = temp
    }
  }
  temp := arr[smallerPartition + 1]
  arr[smallerPartition + 1] = pivot
  arr[end] = temp
  // Return location of pivot.
  return smallerPartition + 1
}
