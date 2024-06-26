package tool

import (
	"strings"
)

func String2IntArray(ids string) []int {
	if ids == "" {
		return make([]int, 0)
	}
	sl := strings.Split(ids, ",")
	arr := make([]int, len(sl))
	for k, v := range sl {
		arr[k] = StringToInt(v)
	}
	return arr
}

func InArrayInt(n int, h []int) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

func InArrayString(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

func Array2String(sl []string, glue string) (ss string) {
	if sl == nil || len(sl) < 1 {
		return ""
	}
	for _, s := range sl {
		if s == "" {
			continue
		}
		if ss == "" {
			ss = s
		} else {
			ss += glue + s
		}
	}
	return ss
}
