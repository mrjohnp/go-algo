package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("******* LEETCODE *******")

	// fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1}))

	// fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))

	// fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))

	// duplicateZeros([]int{0, 4, 1, 0, 0, 8, 0, 0, 3})

	// merge([]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3)

	// removeElement([]int{2, 3, 3, 2, 3, 1, 0, 0, 2, 3, 3, 3, 0}, 3)

	// removeDuplicates([]int{1, 1, 2})

	// checkIfExist([]int{7, 1, 14, 11})

	// validMountainArray([]int{0, 1, 2, 1, 2})

	// replaceElements([]int{65, 1})

	moveZeroes([]int{0, 1, 0, 3, 12})
}

// https://leetcode.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {
	var res int = 0
	var a []int = []int{}

	for i := range nums {
		if nums[i] == 1 {
			a = append(a, nums[i])

			if i == len(nums)-1 && len(a) > res {
				res = len(a)
			}
		} else {
			if len(a) > res {
				res = len(a)
			}

			a = []int{}
		}
	}

	return res
}

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	var res int = 0

	for n := range nums {
		if len(strconv.Itoa(nums[n]))%2 == 0 {
			res++
		}
	}

	return res
}

// https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	for n := range nums {
		nums[n] = nums[n] * nums[n]
	}

	sort.Ints(nums)

	return nums
}

// https://leetcode.com/problems/duplicate-zeros/
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = arr[:len(arr)-1]
			var sub []int = append(arr[:i], 0)
			arr = append(sub, arr[i:]...)
			i++
		}
	}
}

// https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var lastNumIndex int = m - 1

	for i := 0; i < len(nums1) && len(nums2) > 0; i++ {
		if nums1[i] >= nums2[0] || (nums1[i] == 0 && i > lastNumIndex) {
			nums1 = nums1[:len(nums1)-1]
			var sub1 []int = nums1[:i]
			var sub2 []int = nums1[i:]
			sub2 = append([]int{nums2[0]}, sub2...)
			nums1 = append(sub1, sub2...)
			nums2 = nums2[1:]
			lastNumIndex++
		}
	}
}

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}

	return len(nums)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}

	return len(nums)
}

// https://leetcode.com/problems/check-if-n-and-its-double-exist/
func checkIfExist(arr []int) bool {
	var found bool = false

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j]*2 {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	return found
}

// https://leetcode.com/problems/valid-mountain-array/
func validMountainArray(arr []int) bool {
	var res bool = false
	var peak int = 0

	if len(arr) <= 2 || arr[0] == arr[1] || arr[0] > arr[1] {
		res = false
	} else {
		for i := 1; i < len(arr); i++ {
			if arr[i-1] == arr[i] || peak > 1 {
				res = false
				break
			}

			if peak < 1 && arr[i-1] > arr[i] {
				res = true
				peak += 1
				continue
			}

			if peak == 1 && arr[i-1] < arr[i] {
				res = false
				break
			}
		}
	}

	return res
}

// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			arr[len(arr)-1] = -1
		} else {
			var currMax int = 0
			for j := i + 1; j < len(arr); j++ {
				if arr[j] > currMax {
					currMax = arr[j]
				}
			}
			arr[i] = currMax
		}
	}

	return arr
}

// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		if nums[i] == 0 {
			var curr int = 0
			for j := len(nums) - 1; j > i; j-- {
				curr = nums[j]
				fmt.Println("j", curr)
			}
		}
	}
}
