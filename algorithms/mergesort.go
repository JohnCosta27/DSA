package algorithms 

func MergeSort(array []int) []int {
  if (len(array) <= 1) {
    return array
  }

  mid := len(array) / 2
  left := MergeSort(array[0:mid])
  right := MergeSort(array[mid:])

  return mergeLists(left, right)
}

func mergeLists(list1 []int, list2 []int) []int {
  returnList := make([]int, len(list1) + len(list2))
  left := 0
  right := 0

  for (left < len(list1) && right < len(list2)) {
    if (list1[left] < list2[right]) {
      returnList[left + right] = list1[left]
      left++
    } else {
      returnList[left + right] = list2[right]
      right++
    }
  }

  if (left == len(list1)) {
    for (right < len(list2)) {
      returnList[left + right] = list2[right]
      right++
    }
  } else {
    for (left < len(list1)) {
      returnList[left + right] = list1[left]
      left++
    }
  }
  return returnList
}
