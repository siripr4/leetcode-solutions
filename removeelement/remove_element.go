package main

import "fmt"

func main() {
	// fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println("answer: ", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	counter := length
	for index := 0; index < length; {
		fmt.Println("index", index)
		if nums[index] == val {
			counter--
			if index < length-1 {
				fmt.Println("before", nums)
				nums = append(nums[:index], nums[index+1:]...)
				nums = append(nums, -1)
				// nums = nums[index+1:]
				fmt.Println("after", nums)
			} else {
				fmt.Println("before", nums)
				nums = nums[:index]
				nums = append(nums, -1)
				fmt.Println("after", nums)
			}
		} else {
			index++
		}
	}
	fmt.Println(nums)
	return counter
}

func removeElementWithoutSplicing(nums []int, val int) int {
	length := len(nums)
	index := 0

	for index < length {
		if nums[index] == val {
			nums[index] = nums[length-1]
			length--
		} else {
			index++
		}
	}
	return length
}
