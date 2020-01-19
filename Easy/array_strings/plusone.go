/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

*/

func plusOne(nums []int) []int {

	//attempt to optimize on certain inputs
    if first:=nums[len(nums)-1]; first < 9{ // when last index < 9 return
        nums[len(nums) - 1] +=1
        return nums
    }
    
    nines:= true 
    for _, data := range nums{ 
        if data != 9{
            nines = false
        }
    }
    if nines{ // accounting for 'overflow' 
        res:= make([]int, len(nums)+1)
        res[0] = 1
        return res
    }
    carry:= 1
    nums[len(nums) -1] = 0
    for i:= len(nums)-2; i>=0; i--{
        nums[i] += carry
        carry = nums[i] / 10
        nums[i] %= 10
    }
    return nums
}