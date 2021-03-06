package alg

import (
	"fmt"
	"structs"
	"testing"
)

func TestFindMode(t *testing.T) {
	root := structs.TreeNode{
		Val:   1,
		Left:  &structs.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &structs.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	arr := findMode(&root)
	fmt.Println(arr)
}
/**
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
 */
var curNum,max,last int
var res = []int{}
func findMode(root *structs.TreeNode) []int {
	DFS(root)

	return  res
}

func DFS(node *structs.TreeNode){
	if node == nil{
		return
	}
	DFS(node.Left)
	// 中序遍历搜索二叉树相当于遍历有序数组
	if last == node.Val{
		curNum++
	}else {
		// 上一个与这一个不相同，直接变为1
		curNum = 1
	}
	// 比较最大值
	if curNum == max{
		res = append(res,node.Val)
	}else if curNum > max{
		res = []int{node.Val}
		max = curNum
	}
	last = node.Val
	DFS(node.Right)
}
