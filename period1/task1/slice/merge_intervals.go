package main

import (
	"fmt"
	"sort"
)

// 合并重叠区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 1. 按左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	for _, cur := range intervals {
		curL, curR := cur[0], cur[1]
		// 2. 若不重叠，直接加入
		if len(res) == 0 || curL > res[len(res)-1][1] {
			res = append(res, []int{curL, curR})
		} else {
			// 3. 重叠，合并右端点
			last := res[len(res)-1]
			if curR > last[1] {
				last[1] = curR
			}
		}
	}
	return res
}

func main() {
	fmt.Println("合并重叠区间", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
