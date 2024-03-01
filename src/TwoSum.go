package src

type TwoSum struct {
}

func (sum *TwoSum) TwoSum(nums []int, target int) []int {
	var result []int
	count := 0
	for index := 0; index < len(nums); index++ {
		for index2 := 0; index2 < index; index2++ {
			sum := nums[index] + nums[index2]

			if sum == target {
				result = append(result, index2)
				result = append(result, index)
			}
			count++

		}
	}

	return result
}
