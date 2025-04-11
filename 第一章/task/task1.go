// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
package task

import "fmt"

func Singlenumber(nums []int) string {
	resultmap := make(map[int]int)
	// var resultmap map[int]int
	for _, e := range nums {

		resultmap[e]++
	}
	for i := range resultmap {
		if resultmap[i] == 1 {

			fmt.Printf("task1结果:%d\n", i)
			return "============="
		}

	}
	return "task1:非空数组不合题意"

}

// func Rint() {
// 	resultmap := make(map[int]int)
// 	fmt.Printf("%d", resultmap[1])
// }

// resultmap := make(map[int]int)
// 	fmt.Printf("%d", resultmap[1])
