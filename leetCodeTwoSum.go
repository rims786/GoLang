package main

import (
	"fmt"
	"log"
)

// TwoSum finds the indices of the two numbers in the nums slice that add up to the target.
// It returns a slice of the two indices if found, or an error if no solution exists.
func TwoSum(nums []int, target int) ([]int, error) {
	// Create a map to store the number and its corresponding index
	numMap := make(map[int]int)

	// Iterate through the nums slice
	for i, num := range nums {
		// Calculate the complement which is needed to reach the target
		complement := target - num

		// Log the current number and its complement
		log.Printf("Index: %d, Number: %d, Complement: %d", i, num, complement)

		// Check if the complement exists in the map
		if index, found := numMap[complement]; found {
			// Log the found indices
			log.Printf("Found: Index1: %d, Index2: %d", index, i)
			return []int{index, i}, nil
		}

		// Store the number and its index in the map
		numMap[num] = i
		log.Printf("Stored in map: Number: %d, Index: %d", num, i)
	}

	// If no solution is found, return an error
	log.Println("No two sum solution found")
	return nil, fmt.Errorf("no two sum solution")
}

func main() {
	// Example input
	nums := []int{2, 7, 11, 15}
	target := 9

	// Call the TwoSum function
	result, err := TwoSum(nums, target)
	if err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		log.Printf("Result: %v", result)
	}
}
