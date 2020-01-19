/*
Given an unsorted array return whether an increasing 
subsequence of length 3 exists or not in the array.

Formally the function should:

Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.

Input: [5,1,1,2,1,3,1]
Output: true

*/
func increasingTriplet(nums []int) bool {
    if nums == nil || len(nums) < 3{
        return false
    }
    min := nums[0] 
    medium := math.MaxInt32
    
    for i:=1; i< len(nums); i++{
        if nums[i] <= min{ //nums[i] must be the smallest
            min = nums[i]
        } else if  medium != math.MaxInt32 && nums[i] > medium { 
            return true //found our increasing triplet subseq
        } else{ //keep traversing medium
            medium = nums[i]   
        }
    }
    return false
    
}



