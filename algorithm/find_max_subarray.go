package algorithm

// 分治递归法
func FindMaxSubarray(list []float64, start, end int) (int, int, float64) {
	if start == end {
		return start, end, list[start]
	}
	mid := (start + end) / 2
	lLow, lHigh, lSum := FindMaxSubarray(list, start, mid)
	rLow, rHigh, rSum := FindMaxSubarray(list, mid+1, end)
	crossLow, crossHigh, crossSum := findMaxCrossSubarray(list, start, mid, end)
	if lSum >= rSum && lSum >= crossSum {
		return lLow, lHigh, lSum
	} else if lSum <= rSum && rSum >= crossSum {
		return rLow, rHigh, rSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func findMaxCrossSubarray(list []float64, start, mid, end int) (int, int, float64) {
	var lSum, sum float64
	var maxLeft, maxRight int
	lSum = list[mid] - 0.1
	for i := mid; i >= start; i-- {
		sum += list[i]
		if sum > lSum {
			lSum = sum
			maxLeft = i
		}
	}

	rSum := list[mid+1] - 0.1
	sum = 0
	for i := mid + 1; i <= end; i++ {
		sum += list[i]
		if sum > rSum {
			rSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, lSum + rSum
}
