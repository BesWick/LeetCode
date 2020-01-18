/*
Given a sorted array nums, remove the duplicates in-place
 such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array 
in-place with O(1) extra memory.


Input = [1,1,2]
Output= [1,2] 

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
*/

func removeDuplicates(nums []int) int {
    if len(nums) <1{
        return 0
    }
    i:= 0
    for j:=1; j < len(nums); j++{
        if( nums[i] != nums[j]){ //found a unique element with index j
          nums[i+1] = nums[j]
          i++ 
        }
	}
	//get rid of the unmodified latter part of the array
    nums = nums[0:i+1]
    return i+1    
}


