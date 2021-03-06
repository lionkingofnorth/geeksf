/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (52.58%)
 * Likes:    4051
 * Dislikes: 578
 * Total Accepted:    979.1K
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new sorted list. The new
 * list should be made by splicing together the nodes of the first two lists.
 * 
 * Example:
 * 
 * 
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func min(a,b int) int {
	if a<b{return a}else{return b}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{return l2}
	if l2==nil{return l1}
	if l1.Val<l2.Val{
		l1.Next=mergeTwoLists(l1.Next,l2)
		return l1
	}
	l2.Next=mergeTwoLists(l1,l2.Next)
	return l2
}
// @lc code=end

