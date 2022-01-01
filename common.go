package gobmi

import (
	"fmt"
	"strconv"
)

func ConvertStrToFloat64(s string) (parseResult float64, err error) {

	// 将字符串类型转换为float64类型
	parseResult, err = strconv.ParseFloat(s, 64)

	if err != nil {
		err = fmt.Errorf("字符串转换为float64类型失败:%v\n", err)
	}
	return
}

func ConvertFloat64ToStr(f float64) (strResult string) {
	// 将float64类型 取小数点后面精度为2后 转换为字符串。
	strResult = strconv.FormatFloat(f, 'f', 2, 64)
	return
}
