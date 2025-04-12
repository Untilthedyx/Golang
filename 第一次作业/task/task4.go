// 46. 全排列：给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
// 可以使用回溯算法，定义一个函数来进行递归操作，在函数中通过交换数组元素的位置来生成不同的排列，使用 for 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
package task

func Permute(nums []int) [][]int {
	path := make([]int, len(nums))
	used := make([]bool, len(nums))
	var res [][]int
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for s, ok := range used {
			if !ok {
				path[i] = nums[s]
				used[s] = true
				dfs(i + 1)
				used[s] = false
			}
		}
	}
	dfs(0)
	return res
}
