package sorting

import (
	"fmt"
	"sort"
)

func ASC(data sort.Interface) {
	sort.Sort(data)
}

type InputSlice []interface{}

func (sl InputSlice) Len() int {
	return len(sl)
}

func (sl InputSlice) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

func (sl InputSlice) Less(i, j int) bool {
	valI, okI := sl[i].(int)
	valJ, okJ := sl[j].(int)
	if okI && okJ {
		return valI < valJ
	}

	valF, okF := sl[i].(float64)
	valG, okG := sl[j].(float64)
	if okF && okG {
		return valF < valG
	}

	valS, okS := sl[i].(string)
	valT, okT := sl[j].(string)
	if okS && okT {
		return valS < valT
	}

	return fmt.Sprintf("%v", sl[i]) < fmt.Sprintf("%v", sl[j])
}
