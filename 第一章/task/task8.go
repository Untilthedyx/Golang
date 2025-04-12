// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
package task

func Merge(intervals [][]int) [][]int {
	var temp []int
	var j int
	n := len(intervals)
	for i := 0; i < n; i++ {
		for j = 0; j+1 < n-i; j++ {
			if intervals[j][0] > intervals[j+1][0] {
				temp = intervals[j+1]
				intervals[j+1] = intervals[j]
				intervals[j] = temp
			}
		}
	}

	var res [][]int
	res = append(res, append([]int{}, intervals[0]...))
	for r, k := 0, 1; k < n; k++ {
		if res[r][1] < intervals[k][0] {
			res = append(res, append([]int{}, intervals[k]...))
			r++
		} else {
			res[r][1] = max(res[r][1], intervals[k][1])
		}
	}
	return res

}
