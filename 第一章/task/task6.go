// 69. x 的平方根：实现 int sqrt(int x) 函数。
// 计算并返回 x 的平方根，其中 x 是非负整数。
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
// 可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件的平方根或者确定不存在精确的平方根。
package task

func MySqrt(x int) int {
	if x < 0 {
		return -1
	}
	var s, i, j int
	i, j = 0, x
	for i < j {
		s = (i + j) / 2
		if s*s > x {
			j = s - 1
		} else if s*s < x {
			i = s + 1
		} else {
			return s
		}
	}
	if i*i > x {
		return i - 1
	} else {
		return i
	}
}
