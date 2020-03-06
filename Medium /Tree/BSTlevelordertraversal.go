/**
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

Output:

[
  [3],
  [9,20],
  [15,7]
]

*/

//BFS
func levelOrder(root *TreeNode) [][]int {
	 q:= []*TreeNode{}
	 level:= []int{}
	 res := [][]int{}
	 if root == nil{
		 return res
	 }
	 q = append(q, root)
	 level = append(level, 0)
	 for len(q) > 0{
		 topNode:= q[0]
		 topNodeLevel := level[0]
		 //pop top
		 q = q[1:]
		 level = level[1:]
		 if len(res) == topNodeLevel{
			 res = append(res, []int{})
		 }
		 res[topNodeLevel] = append(res[topNodeLevel], topNode.Val) 
		 //add children + level
		 if topNode.Left != nil{
			  q = append(q, topNode.Left)
			 level = append(level, topNodeLevel + 1)
		 }
		 if topNode.Right != nil{
			  q = append(q, topNode.Right)
			 level = append(level, topNodeLevel + 1)
		 }
	 }
	 return res
 }

 //Recursion
 func levelOrder(root *TreeNode) [][]int {
	result :=[][]int{}
	helper(root, 0, &result)
	return result
 }
 
 func helper(root *TreeNode, level int, result *[][]int){
	if root == nil{
		return
	}
	if level == len(*result){
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)

	helper(root.Left, level + 1, result)
	helper(root.Right, level + 1, result)
 }