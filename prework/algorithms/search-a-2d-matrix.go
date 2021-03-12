// https://leetcode.com/problems/search-a-2d-matrix/submissions/
package search

func searchMatrix(matrix [][]int, target int) bool {
	return searchSubMatrix(matrix, target, 0, len(matrix))
}

func searchSubMatrix(matrix [][]int, target int, start int, end int) bool {
	if start >= end {
		return false
	}

	mid := (start + end) / 2
	length := len(matrix[mid])

	if target >= matrix[mid][0] && target <= matrix[mid][length-1] {
		return binarySearch(matrix[mid], target, 0, length)
	} else if target > matrix[mid][length-1] {
		return searchSubMatrix(matrix, target, mid+1, end)
	} else {
		return searchSubMatrix(matrix, target, start, mid)
	}
}

func binarySearch(matrix []int, target int, start int, end int) bool {
	if start >= end {
		return false
	}
	mid := (start + end) / 2

	if target == matrix[mid] {
		return true
	} else if target > matrix[mid] {
		return binarySearch(matrix, target, mid+1, end)
	} else {
		return binarySearch(matrix, target, start, mid)
	}
}
