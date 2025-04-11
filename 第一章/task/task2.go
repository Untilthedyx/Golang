// 198. 打家劫舍：你是一个专业的小偷，计划偷窃沿街的房屋。
// 每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。
// 这道题可以使用动态规划的思想，通过 for 循环遍历数组，利用 if 条件判断来决定是否选择当前房屋进行抢劫，状态转移方程为 dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])。
package task

import "fmt"

func Rob(nums []int) string {

	switch len(nums) {
	case 0:
		fmt.Printf("task2结果：%d\n", 0)
		return "============="
	case 1:
		fmt.Printf("task2结果：%d\n", nums[0])
		return "============="
	case 2:
		fmt.Printf("task2结果：%d\n", max(nums[0], nums[1]))
		return "============="
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])
	dp[2] = max(nums[2]+dp[0], dp[1])
	for i := 3; i < len(nums); i++ {

		dp[i] = max(dp[i-1], dp[i-2]+nums[i], dp[i-3]+nums[i])

	}
	// dp = append(dp, nums[0])
	// dp = append(dp, nums[1])
	// for i := 2; i < len(nums); i++ {

	// 	dp = append(dp, max(dp[i-1], dp[i-2]+nums[i]))

	// }
	fmt.Printf("task2结果：%d\n", dp[len(nums)-1])
	return "============="
}
