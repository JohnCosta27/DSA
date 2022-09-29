package algorithms

func IntMax(ints ...int) int {
	max := ints[0]
	for _, i := range ints {
		if max < i {
			max = i
		}
	}
	return max
}

func FindMaxSubArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	}

	mid := len(arr) / 2
	left := FindMaxSubArray(arr[0:mid])
	right := FindMaxSubArray(arr[mid:])
	midCross := FindBestCrossing(mid, arr)

	return IntMax(left, right, midCross)
}

func FindBestCrossing(mid int, arr []int) int {
	leftSum, leftBestSum := 0, 0
	for i := mid - 1; i >= 0; i-- {
		leftSum += arr[i]
		if leftSum > leftBestSum {
			leftBestSum = leftSum
		}
	}

	rightSum, rightBestSum := 0, 0
	for i := mid + 1; i < len(arr); i++ {
		rightSum += arr[i]
		if rightSum > rightBestSum {
			rightBestSum = rightSum
		}
	}

	return leftBestSum + arr[mid] + rightBestSum
}
