package main

import "fmt"

func maxSubarraySumLessThanTarget(nums []int, target int) int {
	maxSum := 0
	currentSum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		currentSum += nums[end]

		for currentSum >= target && start <= end {
			currentSum -= nums[start]
			start++
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	var n, target int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Printf("Masukkan %d angka dipisah spasi: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Print("Masukkan target: ")
	fmt.Scan(&target)

	result := maxSubarraySumLessThanTarget(nums, target)
	fmt.Println("Output:", result)
}
