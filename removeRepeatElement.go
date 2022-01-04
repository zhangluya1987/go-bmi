package gobmi

import (
	"fmt"
	"sort"
)

func RemoveRepeatElementsForFloat64(oldArray []float64) (newArray []float64) {
	// 如果切片的元素相同 则我们需要进一步做去重处理
	sort.Float64s(oldArray)
	newArray = make([]float64, 0)
	for i := 0; i < len(oldArray); i++ {
		repeatElement := false
		for j := i + 1; j < len(oldArray); j++ {
			if newArray[i] == oldArray[j] {
				repeatElement = true
				break
			}
		}

		if repeatElement == false {
			newArray = append(newArray, oldArray[i])
		}
	}

	//Debug调试
	fmt.Printf("oldArr:%v\n", oldArray)
	fmt.Printf("newArr:%v\n", newArray)
	return
}
