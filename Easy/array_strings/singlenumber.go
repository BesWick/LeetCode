/*

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Ex. Input: [2,2,1]
	Output: 1

	NOTE: Both solution run in O(n) time
*/

//Solution 1:  xor each element in array 
func singleNumberA(nums []int) int {
    result:=0
    for _,data := range nums{
        result^=data
    }
    return result
}



//Solution 2: using hashmap
func singleNumberB(nums []int) int {
    
    mp := make(map[int]int)
    result:= 0
    for _,data := range nums{
        if _,exist := mp[data]; exist{ //if element exist we remove the elemtn
            delete(mp, data)
            
        }else{
            mp[data] = data
        }
	}
	//really O(1) since there's always one element left
    for _, data := range mp{ 
        result = data
    }
    return result 
}