package imath

import (
	"fmt"
	"strconv"
)

// Decimal 保留n位小数
func Decimal(value float64, length int) (res float64) {
	switch length {
	case 1:
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	case 2:
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	case 3:
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	case 4:
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	case 5:
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", value), 64)
	case 6:
		res, _ = strconv.ParseFloat(fmt.Sprintf("%.6f", value), 64)
	default:
		res = 0
	}
	return res
}
