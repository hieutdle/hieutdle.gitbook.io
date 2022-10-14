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

## 9. Palindrome Number

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

```go
121 is a palindrome while 123 is not.
-212 -> 212- is not 
10 -> 01 is not
```

```go
func isPalindrome(x int) bool {
    // When x < 0, x is not a palindrome
    // Also if the last digit of the number is 0, in order to be a palindrome,
    // the first digit of the number also needs to be 0.
    // Only 0 satisfy this property.

    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
	
  revertedNumber := 0;
  for x > revertedNumber {
	  revertedNumber = revertedNumber * 10 + x % 10
	  x /= 10
  }
  // We don't need to revert all the number and compare with the original one.
  // We can cut the number in half and compare  
  // When the length is an odd number
  // we can get rid of the middle digit by revertedNumber/10
  // For example when the input is 12321
  // at the end of the while loop we get x = 12, revertedNumber = 123,
  // since the middle digit doesn't matter in palidrome
  // (it will always equal to itself), we can simply get rid of it.
  return x == revertedNumber || x == revertedNumber/10;
}
```

## 13. Roman to Integer

```go
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

There are six instances where subtraction is used:
* `I` can be placed before `V` (5) and `X` (10) to make 4 and 9.
* `X` can be placed before `L` (50) and `C` (100) to make 40 and 90.
* `C` can be placed before `D` (500) and `M` (1000) to make 400 and 900.

```go
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

```go
func romanToInt(s string) int {
    romanSymbol := map[byte]int{
        'I': 1,
	    'V': 5,
	    'X': 10,
	    'L': 50,
	    'C': 100,
	    'D': 500,
	    'M': 1000,
    }
    l := len(s)
    if l == 0 {
		// If nothing
        return 0
    } else if l == 1{
		// If just one character
        return romanSymbol[s[0]]
    }
    
    sum := romanSymbol[s[l - 1]]
    
    for i := l - 2; i >= 0; i-- {
		// Go from the right to the left. If previous character is smaller => subtract
        if romanSymbol[s[i]] < romanSymbol[s[i+1]] {
            sum -= romanSymbol[s[i]]
        } else {
            sum += romanSymbol[s[i]]
        }
    }
    return sum
}
```

## 1470. Shuffle the Array

Given the array nums consisting of 2n elements in the form `[x1,x2,...,xn,y1,y2,...,yn]`.

Return the array in the form `[x1,y1,x2,y2,...,xn,yn]`.

```go
Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7] 
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].
```
Bad Solution:
```go
func shuffle(nums []int, n int) []int {
    shuffle := make([]int,2*n)
    for i:=0;i<2*n;i++{
        if i % 2 == 0{
            shuffle[i] =  nums[i/2]
        } else {
            shuffle[i] = nums[n+(i/2)]
        }
    }
    return shuffle
}
```
New Solution:
```go
func shuffle(nums []int, n int) []int {
	shuffle := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		shuffle = append(shuffle, nums[i], nums[i+n])
	}
	return shuffle
}
```

## 2114. Maximum Number of Words Found in Sentences

Return the maximum number of words that appear in a single sentence.


```go
Input: sentences = ["please wait", "continue to fight", "continue to win"]
Output: 3
```

```go
func mostWordsFound(sentences []string) int {
    max := 0 
    for _ , v := range sentences {
        count := 1
        for _, v2 := range v{
            if string(v2) == " " {
                count++
            }
        }
        if count > max {
            max = count
        }
    }
    return max
}
```