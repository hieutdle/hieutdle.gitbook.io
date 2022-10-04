# 🥉 Easy


## 1. Two Sum
 Given an array of integers nums and an integer target.
 Return indices of the two numbers such that they add up to target.
 You may assume that each input would have exactly one solution.
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```
```go
func twoSum(nums []int, target int) []int {
        n := len(nums)
	// Check the first element to the second last.
	for i := 0; i < n-1; i++ { // faster than for _,i:= range(nums)
		// Check from the next element to the last element
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
```


