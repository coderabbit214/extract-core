package utils

import "sort"

func InString(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func InInt(target int, strArray []int) bool {
	sort.Ints(strArray)
	index := sort.SearchInts(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// IntersectionStrings 取两个字符串列表的交集
func IntersectionStrings(src []string, dest []string) []string {
	res := make([]string, 0)
	for _, item := range src {
		if InString(item, dest) {
			res = append(res, item)
		}
	}
	return res
}

// IntersectionInts 取两个int列表的交集
func IntersectionInts(src []int, dest []int) []int {
	res := make([]int, 0)
	for _, item := range src {
		if InInt(item, dest) {
			res = append(res, item)
		}
	}
	return res
}

// UnionStrings 取两个字符串列表的并集
func UnionStrings(src []string, dest []string) []string {
	res := make([]string, 0)
	res = append(res, src...)
	for _, item := range dest {
		if !InString(item, res) {
			res = append(res, item)
		}
	}
	return res
}

// UnionInts 取两个int列表的并集
func UnionInts(src []int, dest []int) []int {
	res := make([]int, 0)
	res = append(res, src...)
	for _, item := range dest {
		if !InInt(item, res) {
			res = append(res, item)
		}
	}
	return res
}

func DeleteItem(item string, strArray []string) []string {
	for i, str := range strArray {
		if str == item {
			return append(strArray[:i], strArray[i+1:]...)
		}
	}
	return strArray
}
